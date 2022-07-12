Pods' IP addresses are dynamic and can change over time.

Service is an abstraction with fixed virtual IP address (ClusterIP) fronting a set of pods matching a common label.

Service provides discoverable names and loadbalancing.

Service type defines how the matching Pods are exposed.

* ClusterIP - on cluster-internal IP address (reachable only from within cluster)
* NodePort - on each node's IP address and a static port (30000 - 32767)
* LoadBalancer - on cloud provider's load balancer
* ExternalName - maps a service to a DNS name

A service does not need a deployment but they can work in tandem. A Deployment manages Pods and their replication. A Service routes network requests to a set of Pods.

![svc-deploy](https://user-images.githubusercontent.com/1047259/176018731-d77a9de3-2da7-4b76-9232-915c5e16a1a0.png)

Create deployment + service:

```
kubectl create deployment echoserver --image=k8s.gcr.io/echoserver:1.10 --port=8080 --replicas=5
kubectl expose deployment echoserver --port=80 --target-port=8080
```
