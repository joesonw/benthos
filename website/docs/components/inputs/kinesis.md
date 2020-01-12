---
title: kinesis
type: input
---

```yaml
kinesis:
  batching:
    byte_size: 0
    condition:
      type: static
      static: false
    count: 1
    period: ""
  client_id: benthos_consumer
  commit_period: 1s
  credentials:
    id: ""
    profile: ""
    role: ""
    role_external_id: ""
    secret: ""
    token: ""
  dynamodb_table: ""
  endpoint: ""
  limit: 100
  region: eu-west-1
  shard: "0"
  start_from_oldest: true
  stream: ""
  timeout: 5s
```

Receive messages from a Kinesis stream.

It's possible to use DynamoDB for persisting shard iterators by setting the
table name. Offsets will then be tracked per `client_id` per
`shard_id`. When using this mode you should create a table with
`namespace` as the primary key and `shard_id` as a sort key.

Use the `batching` fields to configure an optional
[batching policy](../batching.md#batch-policy). Any other batching mechanism
will stall with this input due its sequential transaction model.


This input currently provides a single continuous feed of data, and therefore
by default will only utilise a single processing thread and parallel output.
Take a look at the
[pipelines documentation](../pipeline.md#single-consumer-without-buffer) for
guides on how to work around this.

### Credentials

By default Benthos will use a shared credentials file when connecting to AWS
services. It's also possible to set them explicitly at the component level,
allowing you to transfer data across accounts. You can find out more
[in this document](../aws.md).

## Fields

### `batching`

Sorry! This field is currently undocumented.

### `client_id`

Sorry! This field is currently undocumented.

### `commit_period`

Sorry! This field is currently undocumented.

### `credentials`

Sorry! This field is currently undocumented.

### `dynamodb_table`

Sorry! This field is currently undocumented.

### `endpoint`

Sorry! This field is currently undocumented.

### `limit`

Sorry! This field is currently undocumented.

### `region`

Sorry! This field is currently undocumented.

### `shard`

Sorry! This field is currently undocumented.

### `start_from_oldest`

Sorry! This field is currently undocumented.

### `stream`

Sorry! This field is currently undocumented.

### `timeout`

Sorry! This field is currently undocumented.

