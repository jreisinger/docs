# gRPC

* for creating distributed applications and services
* client app can call a method on a server app on a different machine as if it were a local object
* you define a service and its methods (and their parameters and return types) that can be called remotely

![grpc overview](https://grpc.io/img/landing-2.svg)

# Protocol buffers

* by default gRPC uses Protocol Buffers - Google's mechanism for serializing structured data

## 1) Define the structure for data you want to serialize in `.proto` file

* protocol buffer data is structured as messages

```
message Person {
    string name = 1;
    int32 id = 2;
    bool has_ponycopter = 3;
}
```

## 2) Use `protoc` to generate data access classes in your chosen language(s)
