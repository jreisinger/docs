COUPLING is the degree of knowledge one component has to have about another component.

Tightly coupled components

- a change to one component is likely to require corresponding change to the other
- for example both components require the same version of a shared library
- easy to build when optiminizing for the short term

Loosely coupled components

- relatively independent, i.e. easier to upgrade, redeploy or rewrite and even test
- typically interacting via an abstraction (e.g. `io.Writer` instead of `os.File`)
- require more up-front planning

Note: some amount of coupling isn't necessarily a bad thing, especially early
in system's development. One might be tempted to overabstract and
overcomplicate, but premature optimization is the root of all evil.

Note: on occasion, there might be a good reason to tighly couple components.
Abstractions might add overhead that's not acceptable for systems that need
speed.

Note: we'll mostly talk about coupling between services that communicate over a
network, i.e. distributed systems.

MESSAGING PATTERNS

Request-response (synchronous)

- a requester (the client) issues a message and waits for a response from a
  receiver (the server or the service)
- time-based coupling
- e.g. HTTP, REST, RPC, gRPC

```
    c                        s
    l ----- 1) request ----> e
    i                        r
    e                        v
    n <---- 2) response ---- e
    t                        r

```

PUBlish-SUBscribe (asynchronous)

- a requester (the publisher) issues a message, via some kind of messaging
  middleware, which can be retrieved asynchronously and acted on by one or more
  services (the subscribers or consumers)
- producers and consumers are loosely coupled, 
  effectively unaware of each other's existence

```
  publisher1    publisher2
        |          |
        |          |
        v          v
       message/event 
           broker
        ^          ^
        |          |
        |          |
  subscriber1   subscriber2
```

Message

- a piece of data

Event

- signifies that something has happened
- event stream is an ordered and time-stamped sequence of events

Broker

- handles message routing, delivery and persistence
- e.g. Apache Kafka, RabbitMQ, Google Cloud Pub/Sub

Subcriber (consumer) can be

- short-running, often implemented as FaaS, scales horizontally
- long-running, maintains persistent connection to broker and handles stream of
  messages continuosly
