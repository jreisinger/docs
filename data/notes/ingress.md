* the Service object operates at OSI L4 - it only forwards TCP and UDP connections
* if you use Services of `type: NodePort` the clients must connect to a unique port per service
* if you use Services of `type: LoadBalancer` you allocate scarse resources (IP addresses?) for each service
* for HTTP (L7) based services we can do better -> Ingress
* Ingress is k8s' HTTP-based load balancing and "virtual hosting" system
