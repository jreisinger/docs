Nginx deployment

```yaml
# nginx-deploy.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx-deploy
  name: nginx-deploy
spec:
  replicas: 4
  selector:
    matchLabels:
      app: nginx-deploy
  strategy:
    rollingUpdate:      # default, other type is Recreate 
      maxSurge: 2       # maximum above desired number of pods
      maxUnavailable: 1 # maximum unavailable pods
  template:
    metadata:
      labels:
        app: nginx-deploy
    spec:
      containers:
      - image: nginx:1.15
        name: nginx
```

Upgrade nginx version using rolling update strategy

```
$ kubectl edit deployments.apps nginx-deploy # nginx:1.15 => nginx:1.16
$ kubectl rollout status deployment nginx-deploy
Waiting for deployment "nginx-deploy" rollout to finish: 3 of 4 updated replicas are available...
deployment "nginx-deploy" successfully rolled out
```

Rollback

```
$ kubectl rollout undo deployment nginx-deploy
deployment.apps/nginx-deploy rolled back
```

Check revisions

```
$ kubectl rollout history deployment nginx-deploy
deployment.apps/nginx-deploy
REVISION  CHANGE-CAUSE
2         <none>
3         <none>
$ kubectl rollout history deployment nginx-deploy --revision 3
deployment.apps/nginx-deploy with revision #3
Pod Template:
  Labels:	app=nginx-deploy
	pod-template-hash=6ccd96c794
  Containers:
   nginx:
    Image:	nginx:1.15
    Port:	<none>
    Host Port:	<none>
    Environment:	<none>
    Mounts:	<none>
  Volumes:	<none>
```
