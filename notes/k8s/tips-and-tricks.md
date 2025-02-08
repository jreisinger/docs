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
```

Delete pod immediately

```
export now="--force --grace-period=0"
k delete pod nginx $now
```

Explain manifest fields

```
k explain pod.spec.containers.ports [--recursive]
```

Run a command inside a cluster

    k run busybox --image=busybox --rm -it --restart=Never --command -- wget -qO- example.com --timeout 2

Run a shell inside a cluster

    k run alpine --image=alpine --rm -it --restart=Never --command -- /bin/sh
    / # apk --update add bind-tools curl

Copy files

```
k cp <pod>:/path/to/remote/file /path/to/local/file # or vice versa
```

Forward a port from localhost to cluster

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
