# Debugging pods

---

## Deploy demo app

```
k run myapp --image=registry.k8s.io/pause:3.1 --restart=Never
```

---

## Get basic info

```
k get pod $POD
```

---

## Describe details

```
k describe pod $POD
```

---

## Container logs

```
k logs $POD [-c $CONTAINER] 
```

- `--previous` - logs for the previous instance of the container

---

## Exec commands

```
k exec -it $POD -- date
k exec -it $POD -- /bin/sh
```

---

## Add debugging container

```
k debug -it myapp --image=busybox:1.28 --target=myapp
```

---

## Ephemeral containers

* not meant for apps just for debugging
* lack guarantees for resources or execution
* never automatically restarted
* may not have ports
* created using a special `ephemeralcontainers` handler in the API rather than by adding them directly to pod.spec

---

## Debugging via a shell on the node  

```
k debug node/<mynode> -it --image=ubuntu [--profile=sysadmin]
```

- pod name is automatically generated bassed on the node name
- the root filesystem of the Node will be mounted at `/host`
- the container runs in the host IPC, Network, and PID namespaces

---

## More

https://kubernetes.io/docs/tasks/debug/debug-application/debug-running-pod/
