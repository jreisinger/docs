*2017-07-24*

Docker is a container technology a.k.a container runtime. Cointainers standardize software packaging.
It's a well timed fusion of

* kernel features (cgroups, [namespaces](https://gist.github.com/jreisinger/65488e6d7648f3a07a1a346ae3ef549d))
* filesystem tricks (UnionFS)
* networking hacks (iptables)

Think of a container not as a virtual machine but a very lightweight wrapper
around an isolated group of processes. These processes are restricted to private
root filesystem and process namespace.

Docker revision-controls:

1. filesystem layers
2. image tags

# Architecture

<img src="https://raw.github.com/jreisinger/blog/master/files/docker_architecture.png" style="max-width:100%;height:auto;"> 

# Terminology

Docker *server* - the `docker` command run in daemon mode on a Linux host:

```bash
$ sudo docker -d -H unix:///var/run/docker.sock -H tcp://0.0.0.0:2375
```

Docker *container* - a Linux container that has been instantiated from a Docker
image

<img src="https://raw.github.com/jreisinger/blog/master/files/docker_images.png" style="max-width:100%;height:auto;"> 

Docker *image* - one or more filesystem layers and metadata that represent all
the files required to run a Dockerized application

# Images

Two ways to launch a container:

* download a public image
* create your own

To create a custom image you need a `Dockerfile` - each line in a Dockerfile creates a new image layer that is stored by Docker

Build an image:

```bash
git clone https://github.com/spkane/docker-node-hello.git
cd docker-node-hello
docker build -t example/docker-node-hello:latest .
```

Run an image (or a container?):

```bash
docker run -d -p 80:8080 example/docker-node-hello:latest
```

* `-d, --detach` run container in background (daemon mode) and print container ID
* `-p 80:8080` tells Docker to map host's port 80 to the container's port 8080 (port binding)
* `example/docker-node-hello` image to derive the container from
* `:latest` (default) tag specifying the image version

Remove an image:

```bash
docker images
docker rmi <image_id>
```

Remove all images on your Docker host:

```bash
docker rmi $(docker images -q)
```

# Containers

A container is a self-contained execution environment that shares the kernel of
the host system and which is isolated from other containers in the
system.

Containers are a *Linux only* technology.

Create a container (see also "Run an image" above):

```bash
docker run --rm -it ubuntu /bin/bash
```

* `run` = `create` + `start`
* `--rm` - delete the container (the read/write filesystem layer) when it exits
* `-i` - interactive session, i.e. keep STDIN open
* `-t` - allocate a pseudo-TTY
* `/bin/bash` - executable to run within the container

Get into a running container:

```bash
docker ps
docker exec -it <container_id> /bin/bash # new process created in the container
```

Stop a container:

```bash
docker stop <container_id>
```

Remove a container:

```bash
docker ps -a
docker rm <container_id>
```

Remove all containers on your Docker host:

```bash
docker rm  $(docker ps -a -q)
```

# Volumes

* the read/write filesystem layer is a copy-on-write snapshot of an image
* heavy reliance on the read/write filesystem layer isn't the best storage solution (for data intensive apps like DBs)
* the read/write filesystem layer gets removed when the container is removed (`docker rm ...`)
* Docker has the notion of volumes that are maintained separately from the union
    filesystem
* volumes can be shared among containers

Add a volume to a container (`-v`):

```bash
$ docker run -v /data --rm --hostname web --name web -d nginx
$ docker inspect -f '{{ json .Mounts }}' web | jq # note source and destination keys
[
  {
    "Type": "volume",
    "Name": "2d80bc1056787f16b71fb0edced98b3036252083044b1c8db536627c2544a121",
    "Source": "/var/lib/docker/volumes/2d80bc1056787f16b71fb0edced98b3036252083044b1c8db536627c2544a121/_data",
    "Destination": "/data",
    "Driver": "local",
    "Mode": "",
    "RW": true,
    "Propagation": ""
  }
]
```

Add *bind volume* (mount volume on the host and in a container simultaneously):

```bash
$ docker run -v /mnt/data:/data --rm --name web -d nginx
$ docker inspect -f '{{ json .Mounts }}' web
[{"Type":"bind","Source":"/mnt/data","Destination":"/data","Mode":"","RW":true,"Propagation":"rprivate"}]
```

* `Source` on the host won't get removed when you remove the container

Have a data volume container:

```bash
$ docker create -v /mnt/data:/data --name nginx-data nginx          # never runs
$ docker run --volumes-from nginx-data -p80:80 --name web -d nginx
```

# Networks

During installation Docker creates three default networks:

```bash
$ docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
0e07cd43ad1b        bridge              bridge              local
1876373e07e4        host                host                local
e3f087868688        none                null                local
```

1. bridge (virtual switch) is the default --> private namespaced network within the host (which acts as poor man's router); you decide which ports get exposed to the outside world
2. with host networking no separate network namespace is used (`docker run --net host ...`)
3. none is for advanced use cases

<img src="https://raw.github.com/jreisinger/blog/master/files/docker_bridge.png" style="max-width:100%;height:auto;"> 

When you use `-p` Docker creates `iptables` rules that route traffic from the host's public interface on the container's interface on the bridge network.

# DNS

If you want to change DNS servers used when building and running images:

```
$ cat /etc/docker/daemon.json
{
    "dns": ["10.10.10.1", "10.10.10.2"]
}
```

To verify:

```
docker run --rm -it busybox cat /etc/resolv.conf
```

# Monitoring and cleanup

Containers' statistics:

```bash
docker stats [--no-stream]
```

Clean up:

```bash
# you will be asked to confirm
docker system prune                               # containers, networks, images
docker images prune                               # only images

# be careful!
docker volume rm $(docker volume ls -qf dangling=true)  # volumes
```

# Limiting a container's resources

* a container has no resource constraints by default
* Docker provides a way to limit memory, CPU and block IO resources
* your kernel must support Linux capabilities (`docker info | grep WARNING`)

Memory

* if the kernel detects that there is not enough memory, it throws an `Out of Memory Exception` and starts killing processes
* any process is subject to killing (including Docker)
* a process that uses lot of memory but has not been running for long time is a most likely candidate to get killed (see [OOMM](https://www.kernel.org/doc/gorman/html/understand/understand016.html) for more)
* Docker adjusts OOM priority in the Docker daemon so it's less likely to get killed
* the OOM priority on containers is not adjusted so they are more likely to be killed than the Docker daemon

To limit the memory resource to 500 MB and forbid access to swap for a container:

```bash
docker run --rm -it --name mem-muncher --memory=500m --memory-swap=500m mem-muncher
```

See [mem-muncher](https://github.com/jreisinger/mem-muncher) and [presentation](https://gist.github.com/jreisinger/2f87098558d541cdbb7eb30b86163c39) for more.

# Swarm

* [tutorial](https://docs.docker.com/engine/swarm/swarm-tutorial/)
* [service placement](https://docs.docker.com/engine/swarm/services/#control-service-placement) (affinity)
* [monitoring](https://github.com/stefanprodan/swarmprom)
* [setup](https://gist.github.com/jreisinger/a196f3e51e3a7069f7f91665025570cf) a simple cluster (in VirtualBox)

# Sources

* Docker: Up & Running (2015)
* Unix and Linux System Administration Handbook, 5th ed. (2017)
* [Building containers from scratch with Go](https://www.safaribooksonline.com/library/view/building-containers-from/9781491988404/) (2017, video)
