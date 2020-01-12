---
title: redis_streams
type: input
---

```yaml
redis_streams: {}
```

Pulls messages from Redis (v5.0+) streams with the XREADGROUP command. The
`client_id` should be unique for each consumer of a group.

The field `limit` specifies the maximum number of records to be
received per request. When more than one record is returned they are batched and
can be split into individual messages with the `split` processor.

Redis stream entries are key/value pairs, as such it is necessary to specify the
key that contains the body of the message. All other keys/value pairs are saved
as metadata fields.

