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
