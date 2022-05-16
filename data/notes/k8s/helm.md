Helm is a templating engine and package manager for the Kubernetes "operating system".

In Helm's vocabulary, a package is called a *chart*. It contains:

```
$ tree
.
├── Chart.yaml  # describes the chart (name, description, version, author)
├── templates   # K8s manifests with templating directives
│   ├── web-app-pod-template.yaml
│   └── web-app-service-template.yaml
└── values.yaml # default configuration
```

Chart.yaml:

```
apiVersion: 1.0.0
name: web-app
version: 2.5.4
```

web-app-pod-template.yaml:

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
      protocol: TCP
  restartPolicy: Always
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
