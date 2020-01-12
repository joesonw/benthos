---
title: kinesis_balanced
type: input
---

```yaml
kinesis_balanced:
  batching:
    byte_size: 0
    condition:
      type: static
      static: false
    count: 1
    period: ""
  credentials:
    id: ""
    profile: ""
    role: ""
    role_external_id: ""
    secret: ""
    token: ""
  dynamodb_billing_mode: ""
  dynamodb_read_provision: 0
  dynamodb_table: ""
  dynamodb_write_provision: 0
  endpoint: ""
  max_batch_count: 1
  region: eu-west-1
  start_from_oldest: true
  stream: ""
```

BETA: This input is a beta component and is subject to change outside of major
version releases.

Receives messages from a Kinesis stream and automatically balances shards across
consumers.

Messages consumed by this input can be processed in parallel, meaning a single
instance of this input can utilise any number of threads within a
`pipeline` section of a config.

Use the `batching` fields to configure an optional
[batching policy](../batching.md#batch-policy).

### Credentials

By default Benthos will use a shared credentials file when connecting to AWS
services. It's also possible to set them explicitly at the component level,
allowing you to transfer data across accounts. You can find out more
[in this document](../aws.md).

### Metadata

This input adds the following metadata fields to each message:

```text
- kinesis_shard
- kinesis_partition_key
- kinesis_sequence_number
```

You can access these metadata fields using
[function interpolation](../config_interpolation.md#metadata).

## Fields

### `batching`

Sorry! This field is currently undocumented.

### `credentials`

Sorry! This field is currently undocumented.

### `dynamodb_billing_mode`

Sorry! This field is currently undocumented.

### `dynamodb_read_provision`

Sorry! This field is currently undocumented.

### `dynamodb_table`

Sorry! This field is currently undocumented.

### `dynamodb_write_provision`

Sorry! This field is currently undocumented.

### `endpoint`

Sorry! This field is currently undocumented.

### `max_batch_count`

Sorry! This field is currently undocumented.

### `region`

Sorry! This field is currently undocumented.

### `start_from_oldest`

Sorry! This field is currently undocumented.

### `stream`

Sorry! This field is currently undocumented.

