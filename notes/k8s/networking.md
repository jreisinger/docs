Container-to-container communication

* containers in a Pod share the same IP address and ports (like processes on a VM)
* containers can communicate via `localhost` hostname

```yaml
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
    command:
    - /bin/sh
    - -c
    - 'while true; do curl --silent --head localhost; sleep 5; done'
```

```sh
k apply -f multi.yaml
k logs multi -c sidecar -f
```

Pod-to-pod communication

* any pod can reach any other pod

```sh
k run busybox --image=busybox --rm -it --restart=Never -- \
wget $(k get pod multi -o jsonpath={.status.podIP})
```

* Pod IP address is unique across all nodes and namespaces
* achieved by assigning dedicated subnet to each node when registering it

```sh
# show subnets of all cluster nodes
k get nodes -o json | jq -r .items[].spec.podCIDR | sort -V
```

Source: CKA study guide (2021)
