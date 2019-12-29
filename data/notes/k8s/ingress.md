Intro

* the Service object operates at OSI L4 - it only forwards TCP and UDP connections
* if you use Services of `type: NodePort` the clients must connect to a unique port per service
* if you use Services of `type: LoadBalancer` you allocate scarse resources (IP addresses?) for each service
* for HTTP (L7) based services we can do better -> Ingress
* Ingress is k8s' HTTP-based load balancing and "virtual hosting" system

Nginx Ingress Controller

* the most popular generic Ingress controller is probably the open source [NGINX Ingress Controller](https://github.com/kubernetes/ingress-nginx/)
* it reads Ingress objects and merges them into NGINX config file and then signals to the NGINX process to restart
* it has many features and options exposed via annotations
