Pods' IP addresses are dynamic and can change over time.

Service is an abstraction with fixed virtual IP address (ClusterIP) fronting a set of pods matching a common label.

Service provides discoverable names and loadbalancing.

Service type defines how the matching Pods are exposed.

* ClusterIP - on cluster-internal IP address (reachable only from within cluster)
* NodePort - on each node's IP address and a static port
* LoadBalancer - on cloud provider's load balancer
* ExternalName - maps a service to a DNS name

A service does not need a deployment but they can work in tandem. A Deployment manages Pods and their replication. A Service routes network requests to a set of Pods.

![image](https://user-images.githubusercontent.com/1047259/130795064-04865c74-17e9-408a-944b-b8ffcacb0557.png)
