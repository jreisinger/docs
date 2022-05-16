Generate pod manifest

```
k run nginx --image=nginx --port=80 --dry-run=client -o yaml
```

Explain manifest fields

```
k explain pod.spec.containers.ports
```

Run a temporary pod inside a cluster

```
# Shell into it.
k run alpine --image=alpine --rm -it -- sh

# Run a command in it.
k run alpine --image=alpine --rm -it --restart=Never --command -- \
env

k run alpine --image=alpine --rm -it --restart=Never --command -- \
wget -O- https://example.com --timeout 2

k run curl --image=curlimages/curl --rm -it --restart=Never --command -- \
curl -s -o /dev/null -w "%{http_code}" -L https://google.com
```

Copy files

```
k cp <pod>:/path/to/remote/file /path/to/local/file # or vice versa
```

Port forwarding

```
k port-forward <pod> 8080:8080  # tunnel: localhost -> k8s master -> k8s worker node
```

Proxy between localhost and K8s API server

```
$ k proxy                                                             # create proxy
Starting to serve on 127.0.0.1:8001
$ curl localhost:8001/api/v1/namespaces/some-ns/pods/                 # get list of pods
$ curl localhost:8001/api/v1/namespaces/some-ns/services/nginx/proxy/ # access nginx svc
```

Get pod subnets

```
for node in $(k get --no-headers nodes | cut -d' ' -f1); do
  echo -ne "$node\t"
  k get node $node -o json | jq .spec.podCIDR
done
```
