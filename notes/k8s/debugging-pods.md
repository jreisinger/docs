# Debugging pods

---

## Get most important info

```
k get pod $POD
```

---
## Describe details

```
k describe pod $POD
```

---
## Cluster events

```
k events -A
```

- namespaced
- persisted in etcd

---
## See container logs

```
k logs $POD [-c $CONTAINER] 
```

---

## Exec into container

```
k exec -it $POD -- /bin/sh
```

---

## Debug with ephemeral container


* start demo pod
`k run demo --image=registry.k8s.io/pause:3.1 --restart=Never`
* try exec into it
`k exec -it demo -- sh`
* add ephemeral busybox container and attach to it
`k debug -it demo --image=busybox:1.28 --target=demo`
  * `--target` parameter targets the process namespace of another container

---

## Debugging via a shell on the node  

```
k debug node/mynode -it --image=ubuntu
```

- pod name is automatically generated bassed on the node name
- the root filesystem of the Node will be mounted at `/host`
