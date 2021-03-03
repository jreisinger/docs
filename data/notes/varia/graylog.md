# Logs, events and metrics

* log or event - record of a specific thing that happened (e.g. someone authenticated, made a web request, or CPU usage was high for five minutes)
* metrics - usually time based set of numbers that give info about something (e.g. how many authentication requests happened, the number of web requests made, or the CPU usage)

Logs provide more context but are more difficult to store and process because it's more data and it's often unstructured.

# Components

Logs flow:

1. Logs producer (VM, container, application)
2. (OPTIONAL message broker like Kafka - you can do some logs pre-processing here)
3. Input
4. Extractor
5. Stream
6. Pipeline
7. Alert

## Index

* basic unit of storage for data in Elasticsearch

## Input

* defines the method by which Graylog collects logs (e.g. Raw/Plaintext Kafka Plugin, Beats, Syslog UDP)

## [Extractor](https://docs.graylog.org/en/latest/pages/extractors.html)

* tied to Input
* applied on every message that is received by an input
* instructs Graylog how to extract and transform text data into fields that allow you easy filtering and analysis later on
* e.g. extract the HTTP response code, transform it to a numeric field and attach it as http_response_code field to the message

## Stream

* think of it as tagging of incoming messages
* routes messages into categories in real time
* many uses: message categorization, access control, messages parsing and enrichment, ...
* messages can belong to one or more streams

## Pipeline

* tied to Streams
* you can add new field to a log here
* lets you transform and process messages coming from Streams
* allow for: routing, parsing, dropping, blacklisting, modifying and enriching messages as they flow through Graylog

## Alert 

* periodic search that can trigger some notification when a defined condition is met
* composed of: 1. alert condition 2. alert notification

# Configuration

https://docs.graylog.org/en/latest/pages/configuration/server.conf.html

# Searching

* syntax very close to Lucene's
* if you don't specify a field all fields are included in search
* by default all terms or phrases are OR connected
* AND, OR, and NOT are case sensitive

The following characters must be escaped with a backslash:

```
& | : \ / + - ! ( ) { } [ ] ^ " ~ * ?
```

e.g:

```
resource:\/posts\/45326
```


## Full text

All messages that include ssh or login:

```
ssh login
```

## By field

Messages where the field type includes ssh or login:

```
type:(ssh OR login)
```

Messages where the field type includes exact phrase "ssh login":

```
type:"ssh login"
```

Numeric fields support range queries:

```
http_response_code[500 TO 504] # inclusive
http_response_code{400 TO 404} # exclusive
bytes:{0 TO 64]
http_response_code:>=400
http_response_code:(>=400 AND <500)
timestamp:["2019-07-23 09:53:08.175" TO "2019-07-23 09:53:08.575"] # must be UTC
```

## Patterns

Regexes (see [ES regexes syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/query-dsl-regexp-query.html#regexp-syntax) for more):

```
/ethernet[0-9]+/
```

Wildcards - use `?` to replace a single character or `*` to replace zero or more characters:

```
source:exam?ple.*
```

* leading wildcards are disabled to avoid excessive memory consumption

Fuzziness - search for similar terms:

```
ssh login~
source:example.org~
```
