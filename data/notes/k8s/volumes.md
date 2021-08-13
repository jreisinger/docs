# Volumes

A container using the temporary filesystem (default) vs a Volume:

<img src="https://user-images.githubusercontent.com/1047259/129347362-812374d7-3225-4e51-a4de-2ad9d8942fce.png" style="max-width:100%;height:auto;"> 

NOTE: If you need your data to live longer than your pod see [persistent volumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/).

K8s volume abstraction fixes two problems:

1. ephemeral nature of files in Containers (when a Container restarts the files get lost)
2. need to share files between Containers within a Pod

A process in a Container sees a filesystem view composed from their Docker image (mounted at the root of the FS) and volumes (mounted at specified paths within the image).

* volume lifetime == Pod lifetime
* volume is a directory (with some data) that is accessible to the Containers in a Pod
* `pod.spec.volumes` - what volumes to provide for a Pod
* `pod.spec.containers.volumeMounts` - where to mount volumes into Containers

# Types of Volumes

There are many types of volumes. Some of them are listed here.

## [configMap](https://kubernetes.io/docs/concepts/storage/volumes/#configmap)

* a way to inject configuration data into a Pod

```
apiVersion: v1
kind: Pod
metadata:
  name: configmap-pod
spec:
  volumes:
    - name: config-vol
      configMap:
        name: log-config
        items:
          - key: log_level
            path: log_level
  containers:
    - name: test
      image: busybox
      volumeMounts:
        - name: config-vol
          mountPath: /etc/config
```

* All contents stored in `log_level` entry of the CM are mounted into `/etc/config/log_level` file. This file path is derived from `mountPath` + `path`.

## [emtpyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir)

# Resources

* https://kubernetes.io/docs/concepts/storage/volumes
