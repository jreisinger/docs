Intro
-----

- distributed search engine built on top of Apache Lucene (a search engine
  library)
- use JSON over HTTP API (and get back a JSON reply) for indexing, searching
  and managing settings
- clustered by default, you can easily add/remove servers
- data is automatically divided into shards, which get balanced and replicated
  across the available servers in your cluster
- unlike relational database, which stores data in records or rows, ES stores
  data in documents
- relational DB (SQL): row, table, column, database
- ES (noSQL): document, type, field: value, index

Data organization (layout)
--------------------------

- [logical](https://raw.github.com/jreisinger/blog/master/files/es_logical.jpg):
  important for apps
- [physical](https://raw.github.com/jreisinger/blog/master/files/es_physical.jpg):
  important for admins as it determines performance, scalability, and availability

Types
- logical containers for documents (similar to tables which are containers for
  rows)
- the definition of fields in each type is called *mapping*, ex. `name` would
  be mapped as a `string`, but the `geolocation` field under location as a
  special `geo_point` type:
```
{
  "name": "Elasticsearch Denver",
  "organizer": "Lee",
  "location": {
    "name": "Denver, Colorado, USA",
    "geolocation": "39.7392, -104.9847"
  }
}
```

Getting mappings

    curl 'localhost:9200/<index>/_mapping/?pretty'
    
Searching
---------

Show last **10** events:

    curl -XPOST "http://localhost:9200/nxapi/events/_search?pretty=true&sort=date:asc"

More
----

- Elasticsearch in Action (2015)
