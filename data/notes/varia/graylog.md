# Components

Logs flow:

```
Logs producer -> [Kafka (you can do some logs pre-processing)] -> 
Input -> Extractor -> Stream -> Pipeline (e.g. add new field to a log) -> 
Alert
```

Index - basic unit of storage for data in Elasticsearch

## Input

* defines the method by which Graylog collects logs

## [Extractor](https://docs.graylog.org/en/latest/pages/extractors.html)

* allow to instruct Graylog how to extract data from any text in a received message
* work only on text fields (not numeric fields or anythin else)

## Stream

* think of it as tagging of incoming messages
* route messages into categories in real time
* many uses: message categorization, access control, messages parsing and enrichment, ...
* messages can belong to one or more streams

## Pipeline

* tied to Streams
* run rule(s) against specific event
* allow for: routing, parsing, dropping, blacklisting, modifying and enriching messages as they flow through Graylog

## Alert 

* periodic searches that can trigger some notification when a defined condition is met
* are composed of: 1. alert condition 2. alert notification

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
