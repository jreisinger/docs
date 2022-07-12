* Pods' IP addresses are dynamic and can change over time
* services provide discoverable names (using DNS control-plane component) and loadbalancing to a set of Pods
* similar to a Deployment, a Service determines the relevant Pods via labels
* Deployment manages a set of Pods (replicas), Service routes network traffic to a set of Pods

![svc-deploy](https://user-images.githubusercontent.com/1047259/176018731-d77a9de3-2da7-4b76-9232-915c5e16a1a0.png)

```
k create deployment echoserver --image=k8s.gcr.io/echoserver:1.10 --replicas=5
```

Types

* ClusterIP - exposes the Service on a cluster-internal IP (only reachable from within the cluster)

```
k create service clusterip echoserver --tcp=80:8080
k run busybox --image=busybox --rm -it --restart=Never -- wget echoserver
```

* NodePort - exposes the Service on each node's IP address at a static port in the range 30000 - 32767 (reachable also from outside of the cluster)

```
k create service nodeport echoserver --tcp=5005:8080
k run busybox --image=busybox --rm -it --restart=Never -- wget --spider echoserver:5005
ADDR=$(k get nodes -o jsonpath='{ $.items[*].status.addresses[?(@.type=="InternalIP")].address }')
PORT=$(k get svc echoserver -o jsonpath={.spec.ports[0].nodePort})
wget --spider $ADDR:$PORT
```

* LoadBalancer - exposes the Service externally using a cloud provider's load balancer

```
k create service loadbalancer echoserver --tcp=5005:8080
minikube tunnel
k get svc echoserver # see EXTERNAL-IP
```

Source: CKA study guide (2021)
