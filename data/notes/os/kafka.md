Pub/Sub (Publish/Subscribe) messaging system.

* messages organized into *topics* (of interest for consumers)
* *producers* push messages
* *consumers* pull messages
* Kafka runs in a cluster -> nodes are called *brokers*

Message is uniquely identified by:

1. topic
2. partition (like a log, can be kept for some time)
3. offset (like a log line)
