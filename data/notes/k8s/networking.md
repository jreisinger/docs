Container-to-container communication

* share the same IP address and ports
* like processes on a VM (Pod)
* can communicate via `localhost` hostname

```
# multi.yaml
apiVersion: v1
kind: Pod
metadata:
  name: multi
spec:
  containers:
  - name: nginx
    image: nginx
  - name: sidecar
    image: curlimages/curl
    args:
    - /bin/sh
    - -c
    - 'while true; do curl localhost; sleep 5; done'
```

```
k apply -f multi.yaml
k logs multi sidecar -f
```

Pod-to-pod communication

* any pod can reach any other pod

```
k run busybox --image=busybox --rm -it --restart=Never -- wget --spider $(k get pod multi -o jsonpath={.status.podIP})
```

* Pod IP address is unique across all nodes and namespaces
* achieved by assigning dedicated subnet to each node when registering it

```
# show subnets of all cluster nodes
k get nodes -o json | jq -r .items[].spec.podCIDR | sort -V
```
