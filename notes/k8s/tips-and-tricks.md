Container tools

* docker - container runtime + tool for managing containers and images
* containerd - container runtime (slowly replacing docker)
* crictl - CLI for CRI-compatible container runtimes
* podman - tool for managing containers and images

Shell setup

```
source <(kubectl completion bash)
alias k="kubectl"
complete -o default -F __start_kubectl k

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
k run tmp --image=alpine --rm -it --restart=Never -- /bin/sh
/ # apk update && apk add bind-tools curl
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

Show workloads violating a Pod Security Standard

```
k label ns --all pod-security.kubernetes.io/enforce=baseline --overwrite --dry-run=server
```
