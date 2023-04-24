* global service for operational monitoring
* health and performance monitoring

Components

* dashboards
* metrics (free set + paid detailed metrics) and anomaly detection (ML powered)
* alarms -> automatic actions based on thresholds: auto scale, send message via SNS topic, ...
* eventBridge - send events from your app to different services (event-driven architecture)
* logs - central logs repository
* log, container (EKS + ECS), lambda insights - metrics

Metrics
* time-ordered set of data points
* lika a variable that's assuming various values (data points) over time 
* ex: CPU usage of an EC2 instance, latency of an ELB load balancer

Namespaces
* containers (isolation) for metrics
