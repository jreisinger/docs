Useful output flags for `kubectl`:

```sh
-o wide       # more details
-o json       # complete object in JSON format
--v=6         # verbosity
--no-headers
--sort-by=.metadata.creationTimestamp
```

Clean up objecs:

```sh
kubectl delete deployments --all [--selector="app=myapp,env=dev"]
```

Explain resource types:

```sh
kubectl explain svc
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

Create a proxy server between localhost and K8s API server:

```sh
kubectl proxy &                  # create proxy
curl localhost:8001/api/v1/pods  # get list of pods
```

Run containers for troubleshooting:

```
kubectl run alpine    --image=alpine          --rm -it --restart=Never           -- sh

kubectl run nslookup  --image=busybox         --rm -it --restart=Never --command -- nslookup <service>
kubectl run wget      --image=busybox         --rm -it --restart=Never --command -- wget -qO- <service>:<port>
kubectl run curl      --image=curlimages/curl --rm -it --restart=Never --coomand -- curl -LI <service>:<port>
```

* `--command` -- command to run instead of container's default entrypoint

Remove pod stuck in terminating state:

```
kubectl delete pod postgres-86d59f8fb-pjtgx --force --grace-period=0
```

Suspend a cronjob:

```
kubectl patch cronjobs <cronjob> -p '{"spec" : {"suspend" : true }}'
```

* if the cronjob is suspended for too long you get:

```Events:
  Type     Reason            Age                 From                Message
  ----     ------            ----                ----                -------
  Warning  FailedNeedsStart  11s (x6 over 111s)  cronjob-controller  Cannot determine if job needs to be started: too many missed start time (> 100). Set or decrease .spec.startingDeadlineSeconds or check clock skew
```
