* global service for operational (health and performance) monitoring

Components
* *metrics* (free set + paid detailed metrics) and anomaly detection (ML powered)
* central *logs* repository
* alarms - automatic actions based on thresholds: auto scale, send message via SNS topic, ...
* eventBridge - send events from your app to different services (event-driven architecture)

Metrics
* time-ordered set of data points
* like a variable that's assuming various values (data points) over time 
* ex: CPU usage of an EC2 instance, latency of an ELB load balancer

Namespaces
* containers (isolation) for metrics
