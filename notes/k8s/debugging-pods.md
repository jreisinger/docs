# Debugging pods

---

## Deploy demo app

```
k run myapp --image=registry.k8s.io/pause:3.1 --restart=Never
```

---

## Get basic info

```
k get pod myapp
```

---

## Describe details

```
k describe pod myapp
```

---

## Container logs

```
k logs myapp
```

- `-c $CONTAINER` - specify container
- `--previous` - logs for the previous instance of the container

---

## Exec commands

```
k exec -it myapp -- date
k exec -it myapp -- /bin/sh
```

---

## Add ephemeral debug container

```
k debug -it myapp --image=busybox:1.28 --target=myapp
...
/ # ps aux
PID   USER     TIME  COMMAND
    1 root      0:00 /pause
   12 root      0:00 sh
   23 root      0:00 ps aux
```

---

## Ephemeral containers

* not meant for apps just for debugging
* lack guarantees for resources or execution
* never automatically restarted
* cannot have ports
* created using a special `ephemeralcontainers` handler in the API rather than by adding them directly to pod.spec

---

## Debugging via a shell on the node  

```
k debug node/<mynode> -it --image=ubuntu [--profile=sysadmin]
```

- pod name is automatically generated based on the node name
- the root filesystem of the Node will be mounted at `/host`
- the container runs in the host IPC, Network, and PID namespaces

---

## More

https://kubernetes.io/docs/tasks/debug/debug-application/debug-running-pod/
