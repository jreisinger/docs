# Terminology and concepts

Metrics, Traces and Logs -> when, where

* when: all these data types contain creation timestamp
* where: reserved tags (host, device, source, service, env, version) indicate where the data originated from
* metrics -> **what** is happening (e.g. what is the latency, what is number of page views, what is the number of instances in an autoscaling deployment)
* traces -> **how** is sth. happenning (e.g. how are requests  being processed)
* logs -> **why** is sth. happenning (e.g. why did a request fail, why did a user visit a page, why did a cluster autoscale)

Logs

* info from logs can be extracted in the form of attributes
* DD takes into account that not all ingested logs are of equal value
* so log ingestion and indexing are decoupled
* you can collect all logs and
    * archive the ones you don't need (you can Rehydrate them as needed)
    * index and manage the relevant ones

All DD agent and apps configurations in Docker are done via:

* environment variables
* volumes
* Docker labels

Integrations

* bring together all metrics and logs from a source and pass them to DD
* there are intergrations for various sources available

Integration types

* Agent-based
* Authentication (crawler) based: credentials + API (AWS, Slack, PagerDuty)
* Library: library for your language (Node.js, Python) calling Datadog API
* Custom check

Pipelines

* ordered set of processors
* applied to a filtered subset of ingested logs (after collection but before indexing so all logs are processed)
* the resulting logs have a uniform structure with standard attribute names and normalized time/date

Pipeline types

1. OOTB intergration pipelines for common log sources
2. custom pipelines for custom log sources

Processor types

* Grok Parser - extracts attributes from semi-structured text messages
* Several remappers - remap source attributes to standard attributes
* Category processor - enriches logs with attributes that categorize them
* Lookup processor - defines mapping between a log attribute and a human readable value saved in an Enrichment Table or the processors mapping table

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

# Sending logs to DD

1) Install DD agent as Docker container:

```
docker run --cgroupns host --pid host --name dd-agent -v /var/run/docker.sock:/var/run/docker.sock:ro -v /proc/:/host/proc/:ro -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro -e DD_API_KEY="$DD_API_KEY" -e DD_ENV=dev -e DD_TAGS=jr gcr.io/datadoghq/agent:7
```

2) Check Events > Explorer, it should show (filter for `tags:jr`):

```
Datadog agent (v. 7.XX.X) started on <Hostname>
```

3) See Metrics > Explorer (from `jr`).

4) Agent sends only metrics by default, to send also logs add:

* `-e DD_LOGS_ENABLED=true` - enable log collection from the agent container
* `-e DD_LOGS_CONFIG_CONTAINER_COLLECT_ALL=true` - enable log collection from the all other containers
* `-e DD_LOGS_CONFIG_DOCKER_CONTAINER_USE_FILE=true`
* `-e DD_CONTAINER_EXCLUDE="name:datadog-agent"`

5) See Logs > Search (Search for `tags:jr`)

DD [agent logs](https://docs.datadoghq.com/agent/logs) related parameters

* `source` - defines which integration is sending the logs; or it's a custom name
* `service` - the name of the service owning the logs; defaults to the container short-image if no Autodiscovery logs configuration is present (to override add Autodiscovery Docker labels/pod annotations, ex.: `"com.datadoghq.ad.logs"='[{"service": "service-name"}]'`

# Processing logs

There are three ways to work with logs in a unified way:

1. Make sure logs from various sources have the same syntax and naming convention (impossible! :-).
2. Device complicated queries that take into account all relevant logs.
3. Normalize logs into JSON with standard attribute names via pipelines.

JSON logs are parsed automatically and attibutes are extracted.

Semi-structured (non-JSON) logs are parsed via Grok processor.

Attributes - info extracted from logs, e.g.:

<img width="897" alt="image" src="https://user-images.githubusercontent.com/1047259/192524551-b2f6b980-41ed-464a-b6c1-86433a268534.png">

# Web GUI

DD Events

* informational messages about your system that are consumed by events explorer 

Tags

* you can observe aggregate datapoints across several hosts
* bind different data types to allow for correlation between metrics, traces, and logs
* DD recommends looking at containers, VMs, and cloud infra at the `service` level in aggregate (e.g. look at CPU usage across a collection of hosts that represent a service, rather than CPU usage for server A and B)

Log facets

* user-defined tags and attributes from your indexed logs
* meant for qualitative or quantitative data analysis
* once a facet is created, its content is populated for all new logs flowing in the index

Search query can contain

* assigned tags like `env` and `service`
* attributes extracted from the logs like `@http.status_code`
* text strings from log messages
* see [search syntax](https://docs.datadoghq.com/logs/explorer/search_syntax) for more

# Searching logs

All logs without docker_image tag:

```
-docker_image:*
```

# Sources

* https://docs.datadoghq.com
* https://learn.datadoghq.com/courses/take/intro-to-log-management
