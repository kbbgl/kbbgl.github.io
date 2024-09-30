---
slug: data_models_query_langs
title: Data Models and Query Languages
description: Fill me up!
authors: [kbbgl]
tags: [system_design,data_model,query]
---

This document compares relational, document and few graph-based data models. Also, compare some query languages.

## Relational vs Document vs Graph Model

In the relational model, data is organized into relations (tables) where each relation is an unordered collection of tuples (rows).

The relation model use cases were transaction (e.g. sales, reservations) or batch processing (e.g. invoicing, payroll).

If the data in the application has a document-like structure (i.e. a tree of one-to-many relationships where the entire tree is loaded at once), then it's probably better to use the document model. However, if the application does use many-to-many relationships, the relational model is better since the data processing then moves to the database which is likely to have better performance for querying these types of relationships. When any object in the application can have a relationship to any other object, a graph model would be a good selection.

## Object-Relational Mismatch

Most applications today use object-oriented programming. In order for the application to be able to interact with the relational data model, it needs some type of **object-relational mapping (ORM)**.  
