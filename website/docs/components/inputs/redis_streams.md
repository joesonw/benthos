---
title: redis_streams
type: input
---

```yaml
redis_streams:
  body_key: body
  client_id: benthos_consumer
  commit_period: 1s
  consumer_group: benthos_group
  limit: 10
  start_from_oldest: true
  streams:
  - benthos_stream
  timeout: 5s
  url: tcp://localhost:6379
```

Pulls messages from Redis (v5.0+) streams with the XREADGROUP command. The
`client_id` should be unique for each consumer of a group.

The field `limit` specifies the maximum number of records to be
received per request. When more than one record is returned they are batched and
can be split into individual messages with the `split` processor.

Redis stream entries are key/value pairs, as such it is necessary to specify the
key that contains the body of the message. All other keys/value pairs are saved
as metadata fields.

## Fields

### `body_key`

Sorry! This field is currently undocumented.

### `client_id`

Sorry! This field is currently undocumented.

### `commit_period`

Sorry! This field is currently undocumented.

### `consumer_group`

Sorry! This field is currently undocumented.

### `limit`

Sorry! This field is currently undocumented.

### `start_from_oldest`

Sorry! This field is currently undocumented.

### `streams`

Sorry! This field is currently undocumented.

### `timeout`

Sorry! This field is currently undocumented.

### `url`

Sorry! This field is currently undocumented.

