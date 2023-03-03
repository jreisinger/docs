Helm is a templating engine and package manager for the Kubernetes "operating system". It helps you manage application settings and variables. 

In Helm's vocabulary, a package is called a *chart*. It [contains](https://github.com/bmuschko/cka-study-guide/tree/master/ch04/templating-tools/helm):

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

Other two important concepts besides Chart are:

* Repository - place where Charts are collected and shared (like CPAN but for k8s)
* Release - an instance of a Chart running in a K8s

Commands:

```sh
# template locally and display on a console
helm template .

# bundle the template files into a chart archive file (.tgz)
helm package .

# add a Repository
helm search hub wordpress # or https://artifacthub.io
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo list

# install Release with default values
helm install happy-panda bitnami/wordpress

# install another Release with customized values
echo '{mariadb.auth.database: user0db, mariadb.auth.username: user0}' > values.yaml
helm install --generate-name -f values.yaml bitnami/wordpress

# upgrade existing Release
helm upgrade happy-panda -f values.yaml bitnami/wordpress
helm get values happy-panda

# rollback existing Release
helm rollback happy-panda 1

# get status and list of Releases
helm status happy-panda
helm list --all

# uninstall Release
helm uninstall happy-panda
```

Sources:

* https://helm.sh/docs/intro/using_helm/
* Benjamin Muschko: Certified Kubernetes Administrator (CKA) Study Guide (2022)
* Butcher, Farina, Dolitsky: Learning Helm (2021)
