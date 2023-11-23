Reviewed: 2023-09-22

# Terminology and concepts

Logs, metrics and Traces

* logs:     WHY is sth. happenning (e.g. why did a request fail, why did a cluster autoscale)
* metrics:  WHAT is happening (e.g. what is number of page views, what is the number of instances in an autoscaling deployment)
* traces:   HOW is sth. happenning (e.g. how are requests being processed)
* WHEN:     all these data types contain creation timestamp
* WHERE:    three of the reserved attributes (host, source, service) and two of the recommended tags (env, service) indicate where the data originated from

Logs

* info from logs can be extracted in the form of ATTRIBUTES
* DD takes into account that not all ingested logs are of equal value
* so log ingestion and indexing are decoupled
* you can collect all logs and
  * archive the ones you don't need (you can Rehydrate them as needed)
  * index and manage the relevant ones

Integration types

* Agent-based
* Authentication (crawler) based: credentials + API (AWS, Slack, PagerDuty)
* Library: library for your language (Node.js, Python) calling Datadog API
* Custom check

Indexes

* allow you to filter logs into value groups with different retention periods, quotas, usage monitoring and billing
* all logs that match the index filter will flow throught the index
* you can use exclusion filters to exclude some logs
* excluded logs still flow through Live Tail, so they can be still used to generate metrics or can be archived
* all DD organizations have a main index, you can contact Support to have multiple indexes enabled

Archives

* for storing logs you don't need on hand all the time
* store logs in user-provided cloud storage for long-term retention and quick rehydration whenever needed
* after logs are ingested and pass through configured processing pipelines, logs matching the filter query in an Archive are sent to a user-defined cloud storage
* can only be configured by users with admin permissions

Attributes

* info extracted from logs, e.g.:

<img width="897" alt="image" src="https://user-images.githubusercontent.com/1047259/192524551-b2f6b980-41ed-464a-b6c1-86433a268534.png">

* reserved attributes (host, source, status, service, trace_id, message) - automatically generated
* standard attributes - default ones that can be customized
* aliasing to standard attributes - once you have standard attributes settled

Tags

* bind different data types to allow for correlation between logs, metrics and traces
* let you observe aggregate datapoints across several hosts
* DD recommends looking at containers, VMs, and cloud infra at the `service` level in aggregate (e.g. look at CPU usage across a collection of hosts that represent a service, rather than CPU usage for server A and B)

Log facets

* user-defined tags and attributes from your indexed logs
* meant for qualitative or quantitative data analysis
* once a facet is created, its content is populated for all new logs flowing in the index

DD Events

* records of notable changes relevant for managing and troubleshooting IT operations
* e.g. code deployments, service health, configuration changes, or monitoring alerts

# Sending logs to DD

**Note**: When sending logs in a JSON format to Datadog, there is a set of reserved attributes that have a specific meaning within Datadog.

There are many ways you can [send logs](https://docs.datadoghq.com/logs/log_collection) to DD:

- directly to a DD logging endpoint via TCP, TLS or HTTPS
- via log-shipper daemon (rsyslog, syslog-ng, fluentd, ...)
- via DD agent from
    - host
    - application
    - container (agent can collect logs from all the containers)
- via DD forwarder that is an AWS lambda function (serverless)
- via integration

## Via DD agent in Docker container on macOS

1) Install DD agent as Docker container:

```
docker run --name datadog-agent \
--cgroupns host \
--pid host \
-v /var/run/docker.sock:/var/run/docker.sock:ro \
-v /proc/:/host/proc/:ro \
-v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
-e DD_API_KEY="$DD_API_KEY" \
-e DD_ENV=dev \
-e DD_TAGS=jrtest \
gcr.io/datadoghq/agent:7
```

* configuration for Docker DD agent and apps is all done via environment variables (in this ^ case), volumes or Docker labels

2) Check Service Mgmt > Events, it should show (filter for `tags:jrtest`):

```
Datadog agent (v. 7.XX.X) started on <Hostname>
```

3) See Metrics > Explorer (from `jrtest`).

4) Agent sends only metrics by default, to send also logs add:

* `-e DD_LOGS_ENABLED=true` - enable logs collection
* `-e DD_LOGS_CONFIG_CONTAINER_COLLECT_ALL=true` - enable log collection from all other containers
* `-e DD_CONTAINER_EXCLUDE="name:datadog-agent"`

5) See Logs > Search (Search for `tags:jrtest`)

DD [agent logs](https://docs.datadoghq.com/agent/logs) related parameters

* `source` - defines which integration is sending the logs; or it's a custom name
* `service` - the name of the service owning the logs; defaults to the container short-image if no Autodiscovery logs configuration is present (to override add Autodiscovery Docker labels/pod annotations, ex.: `"com.datadoghq.ad.logs"='[{"service": "service-name"}]'`

# Processing logs

JSON logs are parsed automatically and attributes are extracted.

Semi-structured (non-JSON) logs are parsed via Grok processor.

Pipelines

* ordered set of PROCESSORS
* applied to a filtered subset of ingested logs (after collection but before indexing so ALL logs are processed)
* the resulting logs have a uniform structure with standard attribute names and normalized time/date
* pipeline types
  1. OOTB intergration pipelines for common log sources
  2. custom pipelines for custom log sources
* processor types
  * Grok Parser - extracts attributes from semi-structured text messages
  * Several remappers - remap source attributes to standard attributes
  * Category processor - enriches logs with attributes that categorize them
  * Lookup processor - defines mapping between a log attribute and a human readable value saved in an Enrichment Table or the processors mapping table

To have unified log attributes: 

- device complicated queries that take into account all relevant logs -> unwieldy
- use standard attributes
- use [aliasing](https://docs.datadoghq.com/logs/log_configuration/attributes_naming_convention/#aliasing) to standard attributes
- normalize logs into JSON with standard attribute names via pipelines

# Searching logs

Search query can contain

* assigned tags like `env` and `service`
* attributes extracted from the logs like `@http.status_code`
* text strings from log messages
* see [search syntax](https://docs.datadoghq.com/logs/explorer/search_syntax) for more

All logs without `docker_image` tag:

```
-docker_image:*
```

# Sources

* <https://docs.datadoghq.com>
* <https://learn.datadoghq.com/courses/take/intro-to-log-management>
