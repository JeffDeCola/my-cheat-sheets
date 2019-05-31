# RELATIONAL VS NON-RELATIONAL CHEAT SHEET

`relational vs non-relational` _are the two main types of databases._

Documentation and reference,

* A great list of using
  [go with databases](https://github.com/gostor/awesome-go-storage)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## RELATIONAL / SQL / RDBMS

Basically, a relational database is a collection of tables.
Data is displayed in the rows and columns of the table.
Each table has a schema that represents the fixed
attributes and data types that the items in the table will have.

The tables have keys associated with them, which are used to
identify specific columns or rows.

A Structured Query Language (SQL) statements provide
the functionality for managing the database (reading, creating, updating,
and deleting data).

Pros,

* Well-documented and mature.
* Sold and maintained by a number of established corporations.
* SQL standards are well-defined and commonly accepted.
* Very reliable and ACID-compliant.

Cons,

* Schemas rigidly define how all data inserted into the database must be typed and
  composed.
* Don’t work well — or at all — with unstructured or semi-structured data.
* Not suited for large analytics or IoT event loads.
* The tables in your relational database will not necessarily map one-to-one with
  an object or class representing the same data.
* Migration difficult.

Players,

* Oracle
* MySQL
* Microsoft SQL Server
* [PostgreSQL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/postgreSQL-cheat-sheet)
* DB2

## NON-RELATIONAL / NoSQL

NoSQL databases can be schema agnostic, allowing unstructured and semi-structured
data to be stored and manipulated.

Pro,

* Handles massive data.
* Schema-free data models.
* More flexible and easier to administer.
* Generally more horizontally scalable and fault-tolerant.
* Data can easily be distributed across different nodes.

Cons,

* Less widely adopted and mature than RDBMS solutions.
* Not very good for complex data.
* There are a range of formats and constraints specific to each database type.

NoSQL/Non-relational databases can take a variety of forms.

### KEY VALUE STORE

Simple and elegant and crazy fast.
Low-latency data access at any scale.

Players,

* [redis](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/redis-cheat-sheet)
* Amazon DynamoDB

### WIDE COLUMN STORES

Store data in a single column.
More like a multi-dimensional key-value store.
Good for massive scale and massive distributed systems.

Players,

* Cassandra
* Scylla
* HBase

Scylla and Cassandra use an SQL variant called CQL for data definition
and manipulation, making them straightforward to those already familiar with RDBMS.

### DOCUMENT STORES

Store data in the form of JSON documents.
Document stores are similar to key-value or wide column stores,
but the document name is the key and the contents of the document,
whatever they are, are the value.

Players,

* MongoDB
* Couchbase

### GRAPH DATABASES

Represent data as a network of related nodes
or objects in order to facilitate data visualizations
and graph analytics.

Players,

* Neo4J
* Datastax Enterprise Graph

### SEARCH ENGINES

Similar to document stores, but with a greater emphasis
on making your unstructured or semi-structured data easily
accessible via text-based searches.

Players,

* Elasticsearch
* Splunk
* Solr
