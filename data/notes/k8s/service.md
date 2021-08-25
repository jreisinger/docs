Pods' IP addresses are dynamic and can change over time.

Service is an abstraction with fixed virtual IP address (ClusterIP) fronting a set of pods matching a common label.

Service provides discoverable names and loadbalancing.

Service type defines how the matching Pods are exposed.

* ClusterIP - on cluster-internal IP address (reachable only from within cluster)
* NodePort - on each node's IP address and a static port
* LoadBalancer - on cloud provider's load balancer
* ExternalName - maps a service to a DNS name
