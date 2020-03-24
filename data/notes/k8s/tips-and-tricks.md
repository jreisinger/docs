Useful output flags:

```sh
-o wide       # more details
-o json       # complete object in JSON format
-o yaml       # complete object in YAML format
--v=6         # verbosity
--no-headers
```

Sort objects by AGE:

```
kubectl get pod --sort-by=.metadata.creationTimestamp
```

Clean up objecs:

```sh
kubectl delete deployments --all [--selector="app=myapp,env=dev"]
```

Create a proxy server between localhost and K8s API server:

```sh
kubectl proxy &                  # create proxy
curl localhost:8001/api/v1/pods  # get list of pods
```

Explain resource types:

```sh
kubectl explain svc
```

Get pods per node:

```
kubectl get pods -o json --all-namespaces | jq '.items |
  group_by(.spec.nodeName) | map({"nodeName": .[0].spec.nodeName,
  "count": length}) | sort_by(.count) | reverse'
```

Generate resource manifest:

```
kubectl run demo --image=cloudnatived/demo:hello --dry-run -o yaml
```

Show logs:

```sh
kubectl logs [-f] <pod>
kubectl exec -it <pod> -- bash  # or sh instead of bash
```

Copy files:

```sh
kubectl cp <pod>:/path/to/remote/file /path/to/local/file
```

Access Pod via port forwarding:

```sh
kubectl port-forward kuard 8080:8080  # tunnel: localhost -> k8s master -> k8s worker node
```

Run containers for troubleshooting:

```sh
kubectl run demo --image=cloudnatived/demo:hello --expose --port 8888 # pod to troubleshoot
kubectl run nslookup --image=busybox --rm -it --restart=Never --command -- nslookup demo
kubectl run wget --image=busybox --rm -it --restart=Never --command -- wget -qO- http://demo:8888
```

* `--command` -- command to run instead of container's default entrypoint
