# Intro

* the Service object operates at OSI L4 - it only forwards TCP and UDP connections
* if you use Services of `type: NodePort` the clients must connect to a unique port per service
* if you use Services of `type: LoadBalancer` you allocate scarse resources (IP addresses?) for each service
* for HTTP (L7) based services we can do better -> Ingress

# Ingress

* k8s's HTTP-based load balancing and "virtual hosting" system
* at implementation level Ingress is different from pretty much any other k8s resource object
* there is no "standard" Ingress controller built into k8s - you have to pick and install one

## Nginx Ingress Controller

* the most popular generic Ingress controller is probably the open source [NGINX Ingress Controller](https://github.com/kubernetes/ingress-nginx/)
* it reads Ingress objects and merges them into NGINX config file and then signals to the NGINX process to restart
* it has many features and options exposed via annotations
* it parses HTTP request and based on `Host` header and URL path proxies the request to a service

# Manifests

```
# simple-ingress.yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: simple-ingress
spec:
  backend:
    serviceName: my-service
    servicePort: 8080
```
