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
# any HTTP request is forwarded to my-service
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: simple-ingress
spec:
  backend:
    serviceName: my-service
    servicePort: 8080
```

```
# host-ingress.yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: host-ingress
spec:
  rules:
  - host: my-service.example.com
    http:
      paths:
      - backend:
          serviceName: my-service
          servicePort: 8080
```

## TLS

First we need a secret with TLS key and certificate:

```
# tls-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  creationTimestamp: null
  name: tls-secret-name
type: kubernetes.io/tls
data:
  tls.crt: <base64 encoded certificate>
  tls.key: <base64 encoded private key>
```

Then we can reference the secret in Ingress:

```
# tls-ingress.yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: tls-ingress
spec:
  tls:
  - hosts:
    - my-service.example.com
    secretName: tls-secret-name
  rules:
  - host: my-service.example.com
    http:
      paths:
      - backend:
          serviceName: my-service
          servicePort: 8080
```

You can use [cert-manager](https://github.com/jetstack/cert-manager) (and https://letsencrypt.org/) to automate certificates management.

# Sources

* Kubernetes Up & Running (2019)
