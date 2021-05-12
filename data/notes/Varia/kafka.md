Pub/Sub (Publish/Subscribe) messaging system.

* messages organized into *topics* (of interest for consumers)
* *producers* push messages
* *consumers* pull messages
* Kafka runs in a cluster -> nodes are called *brokers*

[Message](https://kafka.apache.org/documentation/#messages) (or event or log or record) is uniquely identified by:

1. topic
2. partition (like a log, can be kept for some time)
3. offset (like a log line)

Kafka is typically used with microservices software architecture (as opposed to monotilith). See [quickstart video](https://kafka.apache.org/quickstart) for more.
