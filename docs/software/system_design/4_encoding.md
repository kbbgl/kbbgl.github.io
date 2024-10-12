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

## Data Flow
