---
title: kafka
type: input
---

```yaml
kafka:
  addresses:
  - localhost:9092
  batching:
    byte_size: 0
    condition:
      type: static
      static: false
    count: 1
    period: ""
  client_id: benthos_kafka_input
  commit_period: 1s
  consumer_group: benthos_consumer_group
  fetch_buffer_cap: 256
  max_processing_period: 100ms
  partition: 0
  sasl:
    enabled: false
    password: ""
    user: ""
  start_from_oldest: true
  target_version: 1.0.0
  tls:
    client_certs: []
    enabled: false
    root_cas_file: ""
    skip_cert_verify: false
  topic: benthos_stream
```

Connects to a kafka (0.8+) server. Offsets are managed within kafka as per the
consumer group (set via config). Only one partition per input is supported, if
you wish to balance partitions across a consumer group look at the
`kafka_balanced` input type instead.

Use the `batching` fields to configure an optional
[batching policy](../batching.md#batch-policy). Any other batching mechanism
will stall with this input due its sequential transaction model.

This input currently provides a single continuous feed of data, and therefore
by default will only utilise a single processing thread and parallel output.
Take a look at the
[pipelines documentation](../pipeline.md#single-consumer-without-buffer) for
guides on how to work around this.

The field `max_processing_period` should be set above the maximum
estimated time taken to process a message.

The target version by default will be the oldest supported, as it is expected
that the server will be backwards compatible. In order to support newer client
features you should increase this version up to the known version of the target
server.

### TLS

Custom TLS settings can be used to override system defaults. This includes
providing a collection of root certificate authorities, providing a list of
client certificates to use for client verification and skipping certificate
verification.

Client certificates can either be added by file or by raw contents:

``` yaml
enabled: true
client_certs:
  - cert_file: ./example.pem
    key_file: ./example.key
  - cert: foo
    key: bar
```

### Metadata

This input adds the following metadata fields to each message:

``` text
- kafka_key
- kafka_topic
- kafka_partition
- kafka_offset
- kafka_lag
- kafka_timestamp_unix
- All existing message headers (version 0.11+)
```

The field `kafka_lag` is the calculated difference between the high
water mark offset of the partition at the time of ingestion and the current
message offset.

You can access these metadata fields using
[function interpolation](../config_interpolation.md#metadata).

## Fields

### `addresses`

Sorry! This field is currently undocumented.

### `batching`

Sorry! This field is currently undocumented.

### `client_id`

Sorry! This field is currently undocumented.

### `commit_period`

Sorry! This field is currently undocumented.

### `consumer_group`

Sorry! This field is currently undocumented.

### `fetch_buffer_cap`

Sorry! This field is currently undocumented.

### `max_processing_period`

Sorry! This field is currently undocumented.

### `partition`

Sorry! This field is currently undocumented.

### `sasl`

Sorry! This field is currently undocumented.

### `start_from_oldest`

Sorry! This field is currently undocumented.

### `target_version`

Sorry! This field is currently undocumented.

### `tls`

Sorry! This field is currently undocumented.

### `topic`

Sorry! This field is currently undocumented.

