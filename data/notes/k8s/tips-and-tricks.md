Alias and completion

```
alias k=kubectl
complete -F __start_kubectl k # enable completion for k alias
```

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
k run alpine --image=alpine --rm -it -- sh
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
$ k proxy                            # create proxy
Starting to serve on 127.0.0.1:8001
$ curl localhost:8001/api/v1/pods    # get list of pods
```

Suspend a cronjob

```
k patch cronjobs <cronjob> -p '{"spec" : {"suspend" : true }}'
```

Remove pod stuck in terminating state

```
k delete pod <pod> --force --grace-period=0
```

Delete objecs by label

```
k delete deployments --all [--selector="app=myapp,env=dev"]
```
