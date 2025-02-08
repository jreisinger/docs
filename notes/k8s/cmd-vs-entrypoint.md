
## Docker

`CMD` defines default command complete with arguments; gets replaced

`ENTRYPOINT` defines default command; gets appended

```
❯ cat Dockerfile
FROM ubuntu
ENTRYPOINT [ "sleep" ]
CMD [ "5" ]
❯ docker build -t my-ubuntu .
❯ docker run my-ubuntu   # sleeps for 5 seconds
❯ docker run my-ubuntu 2 # sleeps for 2 seconds
```

## Kubernetes

`command` corresponds to ENTRYPOINT

`args` corresponds to CMD

```
❯ cat my-ubuntu.yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: my-ubuntu
  name: my-ubuntu
spec:
  containers:
  - image: my-ubuntu
    name: my-ubuntu
    command: [ "sleep" ] # like ENTRYPOINT
    args: [ "5" ]        # like CMD
```

See https://stackoverflow.com/questions/42564058/how-can-i-use-local-docker-images-with-minikube if you want to run the above pod on minikube.