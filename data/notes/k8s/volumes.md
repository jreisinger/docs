A process in a Container sees a filesystem view composed from their Docker image (mounted at the root of the FS) and volumes (mounted at specified paths within the image).

K8s volume abstraction fixes two problems: (1) ephemeral nature of on-disk files in Containers - when it restarts the files get lost (2) need to share files between Containers within a Pod.

* volume lifetime == Pod lifetime
* volume is a directory (with some data) that is accessible to the Containers in a Pod
* `.spec.volumes` - what volumes to provide for a Pod
* `.spec.containers[*].volumeMounts` - where to mount volumes into Containers

There are many types of volumes.

# Types

## configMap

* a way to inject configuration data into a Pod

```
apiVersion: v1
kind: Pod
metadata:
  name: configmap-pod
spec:
  containers:
    - name: test
      image: busybox
      volumeMounts:
        - name: config-vol
          mountPath: /etc/config
  volumes:
    - name: config-vol
      configMap:
        name: log-config
        items:
          - key: log_level
            path: log_level
```

* all contents stored in `log_level` entry of the CM are mounted into `/etc/config/log_level`

# Resources

* https://kubernetes.io/docs/concepts/storage/volumes
