Containers store data in a temporary fielsystem which is empty each time a container starts. A Volume persists data beyond a container restart. A PersistentVolume keeps data even beyond a Pod or node/cluster restart. 

Thus K8s Volume abstraction fixes two problems:

1. Ephemeral nature of files in Containers.
1. Need to share files between Containers within a Pod.

A process in a Container sees a filesystem view composed from their Docker image (mounted at the root of the FS) and Volumes (mounted at specified paths within the image). It can read and write files to this filesystem. A container using the defaut temporary filesystem vs a containers using a Volume:

![129347362-812374d7-3225-4e51-a4de-2ad9d8942fce](https://user-images.githubusercontent.com/1047259/176395171-e934df27-6eac-4704-8a0f-f9c45c6d1d2e.png)

# Types of Volumes

There are many [types of volumes](https://kubernetes.io/docs/concepts/storage/volumes/#volume-types). The type determines the medium backing the volume and its runtime behaviour. Some of the volume types are listed here.

`emtpyDir`

* empty directory within a pod
* read/write access
* only persisted for the lifespan of a pod
* good for cache and data exchange between containers within a pod

```
apiVersion: v1
kind: Pod
metadata:
  name: business-app
spec:
  volumes:
  - name: logs-volume
    emptyDir: {}
  containers:
  - image: nginx
    name: nginx
    volumeMounts:
    - mountPath: /var/logs
      name: logs-volume
```

`hostPath`

* file or directory from the host node's filesystem
* persists data after Pod restart
* presents many security risks

`configMap`

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

`nfs`

* an existing NFS share
* persists data after Pod restart

`persistentVolumeClaim`

* claims a persistent volume

# Persistent Volumes

If you want to persist data even over node and cluster lifetime, like in a database, use Persistent Volumes (PV). The PVC abstracts the underlying implementation details of a PV. A PV can be provisioned statically by creating an object or dynamically by a provisioner assingned to a storage class.

![130926316-1297169b-5d44-442d-a03d-2bcec8468042](https://user-images.githubusercontent.com/1047259/176382749-e72a804f-c3a9-4e05-924b-fcca190e0c84.png)

```
# NOTE: on a kind cluster you don't need to create a PV object manually.
apiVersion: v1
kind: PersistentVolume
metadata:
  name: db-pv
spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce     # RWO, read/write access by a single node
    #- ReadOnlyMany     # ROX, read-only access by many nodes
    #- ReadWriteMany    # RWX, read/write access by many nodes
    #- ReadWriteOncePod # RWOP, read/write access by a single pod
  hostPath:
    path: /data/db
---
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: db-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 512m
---
apiVersion: v1
kind: Pod
metadata:
  name: app-consuming-pvc
spec:
  volumes:
    - name: app-storage
      persistentVolumeClaim:
        claimName: db-pvc
  containers:
  - image: alpine
    name: app
    command: ["/bin/sh"]
    args: ["-c", "while true; do sleep 60; done;"]
    volumeMounts:
      - mountPath: "/mnt/data"
        name: app-storage
```

# Sources

* Certified Kubernetes Administrator (CKA) Study Guide (2022)
