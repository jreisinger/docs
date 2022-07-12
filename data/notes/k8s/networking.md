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

Services

* provide discoverable names (using DNS control-plane component) and loadbalancing to a set of Pods
* similar to a Deployment, a Service determines the relevant Pods via labels
* Deployment manages a set of Pods (replicas), Service routes network traffic to a set of Pods

```
k create deployment echoserver --image=k8s.gcr.io/echoserver:1.10 --replicas=5
```

* `ClusterIP` service type - exposes the Service on a cluster-internal IP (only reachable within the cluster)

```
k create service clusterip echoserver --tcp=80:8080
k run busybox --image=busybox --rm -it --restart=Never -- wget --spider echoserver
```

* `NodePort` service type - exposes the Service on each node's IP address at a static port in the range 30000 - 32767 (reachable from outside of the cluster)

```
k create service nodeport echoserver --tcp=5005:8080
k run busybox --image=busybox --rm -it --restart=Never -- wget --spider echoserver:5005
ADDR=$(k get nodes -o jsonpath='{ $.items[*].status.addresses[?(@.type=="InternalIP")].address }')
PORT=$(k get svc echoserver -o jsonpath={.spec.ports[0].nodePort})
wget --spider $ADDR:$PORT
```

* `LoadBalancer` service type - exposes the Service externally using a cloud provider's load balancer

```
k create service loadbalancer echoserver --tcp=5005:8080
k get svc echoserver # see EXTERNAL-IP
```

Source: CKA study guide (2021)
