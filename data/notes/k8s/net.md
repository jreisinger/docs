Connectivity between containers

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
  - image: nginx
    name: app
    ports:
    - containerPort: 80
  - image: curlimages/curl
    name: sidecar
    args:
    - /bin/sh
    - -c
    - 'while true; do curl localhost; sleep 5; done'
```

Connectivity between pods

* Pod IP address is unique across all nodes and namespaces
* achieved by assigning dedicated subnet to each node when registering it

```
# show subnets of all cluster nodes
k get nodes -o json | jq -r .items[].spec.podCIDR | sort -V
```
