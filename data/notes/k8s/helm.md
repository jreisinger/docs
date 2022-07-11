Helm is a templating engine and package manager for the Kubernetes "operating system".

In Helm's vocabulary, a package is called a *chart*. It contains:

```
$ tree
.
├── Chart.yaml  # describes the chart (name, description, version, author)
├── templates   # K8s manifests with templating directives
│   ├── web-app-pod.yaml
│   └── web-app-service.yaml
└── values.yaml # defaults
```

Chart.yaml:

```
apiVersion: 1.0.0
name: web-app
version: 2.5.4
```

web-app-pod.yaml:

```
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: web-app
  name: web-app
spec:
  containers:
  - image: bmuschko/web-app:1.0.1
    name: web-app
    env:
    - name: DB_HOST
      value: {{ .Values.db_host }}
    - name: DB_USER
      value: {{ .Values.db_user }}
    - name: DB_PASSWORD
      value: {{ .Values.db_password }}
    ports:
    - containerPort: 3000
```

web-app-service.yaml:

```
apiVersion: v1
kind: Service
metadata:
  labels:
    app: web-app-service
  name: web-app-service
spec:
  ports:
  - name: web-app-port
    port: {{ .Values.service_port }}
    protocol: TCP
    targetPort: 3000
  selector:
    app: web-app
  type: NodePort
```

values.yaml:

```
db_host: mysql-service
db_user: root
db_password: password
service_port: 3000
```

Commands:

```
$ helm template . # template locally and display on a console
$ helm package .  # bundle the template files into a chart archive file (.tgz)
```

Sources:

* Benjamin Muschko: Certified Kubernetes Administrator (CKA) Study Guide (2022)
* Butcher, Farina, Dolitsky: Learning Helm (2021)
