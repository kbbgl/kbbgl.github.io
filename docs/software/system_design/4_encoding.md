---
slug: encoding
title: Encoding
description: Chapter 4 from Designing Data Intensive Applications
authors: [kbbgl]
tags: [scalability,system,design,encoding,architecture]
---

When a schema or document structure is changes as part of introducing new features. Encoding data using JSON, XML, Protocol Buffers support systems where old and new data and code need to coexist. These formats are also used for data storage and communication of web services, REST, RPCs and queues.

## Formats

Applications work with data either **in memory** (objects, structs, trees, etc.) or, when writing data to a file or sending it over the network, the data is **encoded** in some self-contained sequence of bytes.

The translation from in-memory to byte sequence is called **encoding/serialization/marshalling**. Going from byte sequence to in-memory is called **decoding/parsing/deserialization/unmarshalling**.

**JSON, XML** are now the standard but are textual, verbose and therefore large. **Binary encoding** libraries such as Thrift and Protocol Buffers (protobuf) can decrease the size. Both require a schema describing the encoded payload and include a code generation tool that takes the schema definition and produces classes that implement the schema in different programming languages.

For example, a JSON payload like this:

```json
{
    "userName": "kbbgl",
    "id": 1337,
    "clubs": ["Boca Juniors", "Arsenal"]
}
```

```protobuf
message Person {
    required string user_name       = 1;
    optional int64  id              = 2;
    repeated string clubs           = 3;
}
```

Apache Avro actually has the best compression.

### Schema Evolution

As schemas change, we need to make sure to maintain compatibility.

- We can make changes to field names but not tags (1, 2, 3).
- We can add new fields to the schema by giving it a new tag. You cannot make it a required field though. It needs to be optional or have a default value.
- We can remove a field only if it's optional and you cannot reuse the same tag number.

## Modes of Data Flow

To send some data to another process which you don't share memory with, such as sending data over a network to write it to a file, we need to encode it as a sequence of bytes. **The most common ways for how data flows between processes is using databases, service calls and async message passing**.

In case of databases, the process that writes to the database encodes the data and the process that reads from the database decodes it.

In case of service calls, the most common way is to use clients and servers using REST. Services expose an application-specific API that only allows outputs and inputs predetermined by the business logic.

RPCs are also popular and work on the idea that network requests and local function calls are similar (which they are not). gRPC is a popular RPC and builds on the Protocol Buffers implementation. gRPC supports **streams**, a call consists of not just one request/response but a series of them over time. RPCs are mostly used for communication within microservices running on the same system network.

**Asynchronous message-passing systems** are somewhere between databases and RPCs.
A client request (a **message**) is delivered to another process passing through a **message broker/queue**. The sender doesn't wait for the message to be delivered.
Message brokers have one process, the **publisher**, send a message to a name queue or topic, and the broker ensures that the message is delivered to one or more **consumers/subscribers**.

Using a message queue is more reliable than RPCs.

The **actor model** is a design pattern for concurrency in a single process. Actors are usually an instance of the system which has some logic and state within it and it communicates with other actors by sending/receiving async messages. In this case, there's no need to encode and decode the messages since it's an IPC.
There's also the **distributed actor framework** where the actors are spread across different nodes so the messages are encoded/decoded before and after the message is sent over the network. 
