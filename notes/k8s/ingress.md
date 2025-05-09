# Intro

* the Service object operates at OSI L4 - it only forwards TCP and UDP connections
```
$ kubectl get svc cportal
NAME      TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
cportal   ClusterIP   10.99.61.122   <none>        8080/TCP   79d
```
* if you use Services of `type: NodePort` the clients must connect to a unique port per service
* if you use Services of `type: LoadBalancer` you allocate scarce resources (IP addresses) for each service
* for HTTP (L7) based services we can do better -> Ingress

# Ingress

![image](https://user-images.githubusercontent.com/1047259/178455570-23399dd8-f100-4a45-816e-614bf6a493b2.png)

* k8s's HTTP-based load balancing and "virtual hosting" system
* at implementation level Ingress is different from pretty much any other k8s resource object
* in particular it is split into 1) common resource specification (the Ingress object) and 2) a controller implementation
* there is no "standard" Ingress controller built into k8s - you have to pick and install one

## Nginx Ingress Controller

* the most popular generic Ingress controller is probably the open source [NGINX Ingress Controller](https://github.com/kubernetes/ingress-nginx/)
* it reads Ingress objects and merges them into NGINX config file and then signals to the NGINX process to restart
* it has many features and options exposed via annotations
* it parses HTTP request and based on `Host` header and URL path proxies the request to a service
* it's a pod running NGINX process

```
kubectl get pods -A -l app.kubernetes.io/name=ingress-nginx
```

# Manifests

```
$ k create ingress my-service --dry-run=client -o yaml --rule="my-service.example.com/*=my-service:8080"
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  creationTimestamp: null
  name: my-service
spec:
  rules:
  - host: my-service.example.com
    http:
      paths:
      - backend:
          service:
            name: my-service
            port:
              number: 8080
        path: /
        pathType: Prefix # remove * to get Exact
status:
  loadBalancer: {}
```

## TLS

First we need a secret with private key and certificate for a given host (my-service.example.com in this case):

```
# tls-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  creationTimestamp: null
  name: tls-secret-name
type: kubernetes.io/tls
data:
  tls.key: <base64 encoded private key>
  tls.crt: <base64 encoded certificate>
```

Then we can reference the secret in Ingress via `secretName` field:

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

You can use [cert-manager](https://cert-manager.io/docs/) and e.g. [Let's Encrypt](https://letsencrypt.org/) to automate certificates management.

# Tips and tricks

If you specify duplicate or conflicting configurations of Ingress object, the behavior is undefined.

An Ingress object can only refer to an upstream (backend) service in the same namespace. However, multiple Ingress objects in different namespaces can specify subpaths for the same host. These specifications are then merged together. This means that Ingress needs to be coordinated globally across the cluster.

Get [nginx ingress](https://kubernetes.github.io/ingress-nginx/troubleshooting/) logs

```
kubectl get po -A | grep ingress
kubectl logs -n <namespace> nginx-ingress-controller-67956bf89d-fv58j
```

# Deploy app exposed outside minikube cluster via https

```sh
alias k='kubectl'

minikube start

# Install ingress controller.
k apply -f https://projectcontour.io/quickstart/contour.yaml
k get svc envoy -n projectcontour # EXTERNAL-IP should be <pending>

# Create route to LoadBalancer services.
minikube tunnel
k get svc envoy -n projectcontour # EXTERNAL-IP should be 127.0.0.1

# Create deployment and service.
k create deployment alpaca --image=gcr.io/kuar-demo/kuard-amd64:green --replicas=3 --port=8080
k expose deployment alpaca

# Add following to /etc/hosts:
127.0.0.1   alpaca.example.com

# Create TLS key and cert for alpaca.example.com.
mkcert alpaca.example.com

# Create secret holding TLS key and cert.
k create secret tls alpaca-tls --cert alpaca.example.com.pem --key alpaca.example.com-key.pem

# Create ingress object to expose service outside cluster.
cat <<EOF | k apply -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: tls-ingress
spec:
  tls:
  - hosts:
    - alpaca.example.com
    secretName: alpaca-tls
  rules:
  - host: alpaca.example.com
    http:
      paths:
      - pathType: Prefix
        path: "/"
        backend:
          service:
            name: alpaca
            port:
              number: 8080
EOF

# Install the local CA in the system trust store.
mkcert -install

# Restart Firefox and visit https://alpaca.example.com
```

```sh
# Cleanup.
mkcert -uninstall
minikube delete

# Remove alpaca.example.com from /etc/hosts
```

# Sources

* Kubernetes Up & Running (2nd 2019, 3rd 2022)
