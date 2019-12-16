Some of the [quotes](https://quotes.reisinge.net/) I like best ATM :-)

> The programmer, like the poet, works only sligtly removed from the pure thought-stuff. Yet the program construct, unlike the poet's words, is real in the sense that it moves and works, producing visible outputs separate from the construct itself. -- Frederick P. Brooks

> Money is like gas in the car - you need to pay attention or you'll end up on the side of the road - but a successful business or a well-lived life is not a tour of gas stations. -- Tim O'Reilly, WTF

---

# Components

Streams

* think of it as tagging of incoming meesages
* route messages into categories in real time
* many uses: message categorization, access control, messages parsing and enrichment, ...
* messages can belong to one or more streams

Pipelines

* run rule(s) against specific event
* tied to Streams
* allow for: routing, parsing, dropping, blacklisting, modifying and enriching messages as they flow through Graylog

Alerts are composed of:

1. alert condition
1. alert notification

Index

* basic unit of storage for data in Elasticsearch

Input

* defines the method by which Graylog collects logs

# Searching

* syntax very close to Lucene's
* if you don't specify a field all fields are included in search
* by default all terms or phrases are OR connected
* AND, OR, and NOT are case sensitive

All messages that include ssh or login:

```
ssh login
```

Messages where the field type includes ssh or login:

```
type:(ssh OR login)
```

Messages where the field type includes exact phrase "ssh login":

```
type:"ssh login"
```

Regexes (see [ES regexes syntax](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/query-dsl-regexp-query.html#regexp-syntax) for more):

```
/ethernet[0-9]+/
```

Wildcards - use `?` to replace a single character or `*` to replace zero or more characters:

```
source:exam?ple.*
```

Leading wildcards are disabled to avoid excessive memory consumption.

Fuzziness - search for similar terms

```
ssh logni~
source:example.org~
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
