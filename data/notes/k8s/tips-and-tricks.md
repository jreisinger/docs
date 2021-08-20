```
alias k=kubectl
complete -F __start_kubectl k # enable completion for k alias
```

Output flags

```
-o wide       # more details
-o yaml       # complete object in YAML format
--v=6         # verbosity
--no-headers
--sort-by=.metadata.creationTimestamp
```

Explain resource fields

```
k explain pod.spec.containers.ports.protocol
```

Generate pod manifest

```
k run nginx --image=nginx --dry-run=client -o yaml
```

Troubleshoot

```
k logs [-f] <pod>
k exec -it <pod> -- sh
```

Run a temporary pod inside a cluster and start a shell/command in it

```
k run alpine --image=alpine --rm -it -- sh
k run busybox --image=busybox --rm -it -- nslookup <service>
```

* `--command` -- command to run instead of container's default entrypoint

Copy files

```
k cp <pod>:/path/to/remote/file /path/to/local/file # or vice versa
```

Port forwarding

```
k port-forward kuard 8080:8080  # tunnel: localhost -> k8s master -> k8s worker node
```

Proxy server between localhost and K8s API server

```
k proxy &                  # create proxy
curl localhost:8001/api/v1/pods  # get list of pods
```

Suspend a cronjob

```
k patch cronjobs <cronjob> -p '{"spec" : {"suspend" : true }}'
```

* if the cronjob is suspended for too long you get:

```Events:
  Type     Reason            Age                 From                Message
  ----     ------            ----                ----                -------
  Warning  FailedNeedsStart  11s (x6 over 111s)  cronjob-controller  Cannot determine if job needs to be started: too many missed start time (> 100). Set or decrease .spec.startingDeadlineSeconds or check clock skew
```

Remove pod stuck in terminating state

```
k delete pod <pod> --force --grace-period=0
```

Delete objecs by label

```
k delete deployments --all [--selector="app=myapp,env=dev"]
```
