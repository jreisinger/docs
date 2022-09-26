Pub/Sub messaging system.

* messages (or events or logs or records) are organized into *topics* (of interest for consumers)
* *producers* publish (write) stream of messages
* *consumers* subscribe to (read) stream of messages
* Kafka runs in a cluster where the cluster nodes are called *brokers*

[Message](https://kafka.apache.org/documentation/#messages) is uniquely identified by:

1. topic
2. partition (P) - like a log file, can be kept for some time
3. offset - like a log line

<img width="492" alt="image" src="https://user-images.githubusercontent.com/1047259/151317569-4d669f92-8414-4fa6-9c32-d7f9c3635ad2.png">

Kafka is typically used with microservices software architecture (as opposed to monotilith). See [quickstart video](https://kafka.apache.org/quickstart) for more.

More: https://kafka.apache.org
