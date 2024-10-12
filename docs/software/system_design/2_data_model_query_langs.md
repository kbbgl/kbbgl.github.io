---
slug: data-models-query-langs
title: Data Models and Query Languages
description: Chapter 2 from Designing Data Intensive Applications
authors: [kbbgl]
tags: [design,system_design,data_model,query]
---

If the application has **many-to-many** relationships, better to **use a relational database**.

If the application has mostly **one-to-many** relationships such as tree-structured data, the NoSQL/document data model is preferable.
ORMs (object-relation model) is a layer that translates the relation model to the application OOP.

In case there is a need for `JOIN`ing in a document model, the application will neeed to query multiple documents and merge the documents. This is not optimal since it will affect performance since it moves some of the querying from the database (which is faster for I/O) to the application code.

## Graph Models

If the application has many-to-many relationships but the data is too complex for a relational model, it becomes more natural to start modeling your data as a graph.

A graph consists of **vertices** (aka nodes, entities) and **edges** (aka relationships, arcs). Graph stores can be seen as a relational database with 2 tables, one for vertices and one with edges.

Typically, graph models are used for examples such as:

* Social graphs where the vertices are people and edges represent acquaintenanceship.

* Websites where the vertices are web pages and the edges indicate HTML links to other pages.

* Roads where vertices are junctions and edges represent roads between them.

Graph models can also be used where the vertices are completely different types of objects in a single datastore. For example, Facebook's graph has vertices represent people, locations, events, checkins and comments made by users where edges indicate which people are friends with each other, which checkin happened in which location, who commented on which post, etc.

### Property Graphs

Property graphs can be used in Neo4j for example.

Each vertex consists of:

* A unique identifier
* A set of outgoing edges
* A set of incoming edges
* A collection of properties

Each edge consists of:

* A unique identifier
* The vertex at which the edge starts (tail vertex)
* The vertex at which the edge ends (the head vertex)
* A label to describe the kind of relationship between the two vertices
* A collection of properties (key-value pairs)

Some advantages of using a graph model are:

* The properties allow you to add some anecdotal information about the edge or vertex
* The head and tail vertices allow you to traverse the graph easily.
* Querying the graph for outgoing or incoming edges is easy.
