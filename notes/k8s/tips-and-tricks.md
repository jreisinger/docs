Container tools

* docker - container runtime + tool for managing containers and images
* containerd - container runtime (slowly replacing docker)
* crictl - CLI for CRI-compatible container runtimes
* podman - tool for managing containers and images

Shell setup

```
alias k="kubectl"
export do="--dry-run=client -o yaml"
export now="--force --grace-period=0"
```

Generate pod manifest

```
k run nginx --image=nginx --port=80 $do
```

Explain manifest fields

```
k explain pod.spec.containers.ports
```

Run a temporary pod inside a cluster

```
# Shell into it.
k run tmp --image=alpine --rm -it --restart=Never -- sh
/ # apk update && apk add bash && bash
```

```
# Run a command in it.
k run tmp --image=busybox --rm -it --restart=Never -- wget example.com --timeout 2
k run tmp --image=curlimages/curl --rm -it --restart=Never -- curl example.com --max-time 2
```

Copy files

```
k cp <pod>:/path/to/remote/file /path/to/local/file # or vice versa
```

Port forwarding

```
k port-forward <pod> 8080:8080 # tunnel: localhost -> k8s master -> k8s worker node
```

Get pod subnets assigned to the nodes

```
for node in $(k get --no-headers nodes | cut -d' ' -f1); do
  echo -ne "$node\t"
  k get node $node -o json | jq .spec.podCIDR
done
```
