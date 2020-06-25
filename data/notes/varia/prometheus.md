# Monitoring

Prometheus is a metrics-based tool for operational monitoring of computer systems, i.e. it's useful for:

* alerting - let me know when things go wrong
* debugging - let me investigate why things went wrong
* trending - let me see how systems are being used and changed over time

At the end of the day all moniroting systems are data processing pipelines.

Monitoring is about events like receviving a HTTP request, entering a function, a user logging in, requesting more memory from the kernel. All events also have context like the IP address the request is coming from or call stacks of functions.

Having all context all the time is impractical - ways to reduce the amount of data to something workable:

* profiling - let's have some of the context for limited periods of time (`tcpdump` for networking, eBPF for Linux kernel)
* tracing - let's take some proportion of events (sampling)
* logging - let's have a limited set of events and some of the context for each of these events (Graylog, Splunk)
* metrics - let's track aggregations over time of different types of events largely ingoring context (Prometheus)

# Architecture and Components

```
                      +-------------+
                      |             |
                      | EC2, K8s,   |
                      | Consul, etc.|
                      +------^------+
                             |
                     +--------------------------------+
 +-------------+     |       |                        |
 | Application |     | +-----------+       Prometheus |
 |             |     | |           |                  |
 |    +--------+     | | Service   |                  |
 |    |Client  |     | | Discovery |                  |  +--------------+
 |    |Librabry<---+ | |           |                  |  | Email, PD,   |
 +-------------+   | | +-----------+                  |  | Slack, etc.  |
                   | |       |                        |  +------^-------+
                   | |       |                        |         |
                   | | +-----v -----      +---------+ |  +--------------+
                   | | |           |      |         | |  |              |
 +-------------+   +---+ Scraping  |      | Rules & +----> Alertmanager |
 |  Exporter   <---+ | |           |      | alerts  | |  |              |
 +-------------+     | +-----------+      +--^------+ |  +--------------+
        |            |       |               |   |    |
        |            | +-----v-------------------v--+ |
 +------v------+     | |                            | |  +--------------+
 |             |     | |          Storage           +----> Dashboards   |
 | 3rd Party   |     | |                            | |  +--------------+
 | Application |     | |                            | |
 |             |     | +----------------------------+ |
 +-------------+     |                                |
                     +--------------------------------+

```

## Node exporter

* is a kind of exporter
* exposes kernel- and machine-level metrics on Unix systems (CPU, memory, disk space, disk I/O, network bandwidth, ...)
* *no* metrics about individual processes (in the Prometheus architecture you monitor applications and services directly)

# PromQL

* functional (not an SQL-like) query language for selecting and aggregating time series in real time
* query expression result can be shown as a graph, tabular data or consumed via API by external systems
* PromQL is powerful but most of the time your needs will be simple

https://prometheus.io/docs/prometheus/latest/querying/basics/

If you enter `up` query into the expression browser and hit "Execute" you get:

```
up{instance="localhost:9090",job="prometheus"}          1
```

* `up` - special metric added by Prometheus during scrape (HTTP request)
* `instance` - (default) label indicating the target that was scraped
* `job` - (default) label indicating the type of the application (it comes from `job_name` defined in `prometheus.yaml`)
* `1` - the metric's value (target is up)

## Metrics Types and Aggregations

### Gauge

* snapshot of state, current absolute value
* we usually agregate it with `sum`, `avg`, `min` or `max`

Total FS size on each machine (node_filesystem_size_bytes metric comes from Node exporter):

```
sum(node_filesystem_size_bytes) without(device, fstype, mountpoint)
```

  * sum up everything with the same labels ingnore those three

### Counter

* tracks the number or size of events (total since start)
* use `rate` function as counters are always increasing

How many samples Prometheus is ingesting per-second *averaged* over one minute:

```
rate(prometheus_tsdb_head_samples_appended_total[1m])
```

The output of `rate` is a gauge, so e.g. to get total bytes received per machine per second:

```
sum(rate(node_network_receive_bytes_total[5m])) without(device)
```

## Selectors

You almost always will want to limit by job label (defines application type), e.g.:

```
process_resident_memory_bytes{job="kubelet"}
```

* instant vector (one dimensional list) selector

### Matchers

```
=  --> job="node"
!=
=~ --> jon=~"n.*"  # fully anchored, RE2
!~
```

## Query examples

https://prometheus.io/docs/prometheus/latest/querying/examples/

# Labels

* key-value pairs associated with time series
* together with with metric name uniquely idendify metrics (time series)

There are two types of labels although you don't see any difference among them
in PromQL:

1) Instrumentation labels

* things that are know inside you application

2) Target labels

* identify a specific monitoring target (a target that monitoring scrapes)
* relate more to your architecture
* attached by Prometheus as part of process of scraping metrics
* come from service discovery (metadata) and relabelling
* service discovery metadata are converted to target labels
* default ones: `instance`, `job`

# Kubernetes

You can run Prometheus in K8s and monitor K8s objects in two ways:

1. Standard K8s objects (`kind`s) like configMap, deployment and service + access permissions so Prometheus can access (monitor) K8s objects (sample [manifest](https://raw.githubusercontent.com/prometheus-up-and-running/examples/master/9/prometheus-deployment.yml)).
2. [Prometheus Operator](https://github.com/coreos/prometheus-operator) which uses custom resource definition (CRD) feature of k8s to define custom K8s objects (like Prometheus and PrometheusRule).

P8s can discover targets to monitor by using K8s API. There are currently these types of [K8s service discovery](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#kubernetes_sd_config) you can use with P8s:

* node - discover the nodes the K8s cluster (to monitor the infra around K8s)
* service - useful for monitoring the infrastructure of and under K8s and for blackbox monitoring, to check if the service is responding at all
* endpoints - will return a target for every port of every pod backing each of your services
* pod - will return a target for each port of every one of your pods (even of those not backing any service)
* ingress - you should only use this role for blackbox monitoring

# Tips and tricks

* Prometheus uses base units (such as bytes and seconds) and leaves pretty printing to frontend tools like Grafana
* all Prometheus components run happily in containers, except for the Node exporter
* Prometheus is implemented as a single binary `prometheus` (written in Go)
* `prometheus.yml` is the Prometheus configuration file

Useful metrics

* process_resident_memory_bytes - how much memory is the application using (in bytes)

# Sources

* Prometheus: Up & Running (2018)
