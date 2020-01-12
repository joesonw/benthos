---
title: kafka
type: output
---

```yaml
kafka:
  ack_replicas: false
  addresses:
  - localhost:9092
  backoff:
    initial_interval: 0s
    max_elapsed_time: 5s
    max_interval: 1s
  batching:
    byte_size: 0
    condition:
      type: static
      static: false
    count: 1
    period: ""
  client_id: benthos_kafka_output
  compression: none
  key: ""
  max_in_flight: 1
  max_msg_bytes: 1000000
  max_retries: 0
  partitioner: fnv1a_hash
  sasl:
    enabled: false
    password: ""
    user: ""
  target_version: 1.0.0
  timeout: 5s
  tls:
    client_certs: []
    enabled: false
    root_cas_file: ""
    skip_cert_verify: false
  topic: benthos_stream
```

The kafka output type writes a batch of messages to Kafka brokers and waits for
acknowledgement before propagating it back to the input. The config field
`ack_replicas` determines whether we wait for acknowledgement from all
replicas or just a single broker.

It is possible to specify a compression codec to use out of the following
options: `none`, `snappy`, `lz4` and `gzip`.

Both the `key` and `topic` fields can be dynamically set using
function interpolations described [here](../config_interpolation.md#functions).
When sending batched messages these interpolations are performed per message
part.

The `partitioner` field determines how messages are delegated to a
partition. Options are `fnv1a_hash`, `murmur2_hash`, `random` and
`round_robin`. When a hash partitioner is selected but a message key
is empty then a random partition is chosen.

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

## Fields

### `ack_replicas`

Sorry! This field is currently undocumented.

### `addresses`

Sorry! This field is currently undocumented.

### `backoff`

Sorry! This field is currently undocumented.

### `batching`

Sorry! This field is currently undocumented.

### `client_id`

Sorry! This field is currently undocumented.

### `compression`

Sorry! This field is currently undocumented.

### `key`

Sorry! This field is currently undocumented.

### `max_in_flight`

Sorry! This field is currently undocumented.

### `max_msg_bytes`

Sorry! This field is currently undocumented.

### `max_retries`

Sorry! This field is currently undocumented.

### `partitioner`

Sorry! This field is currently undocumented.

### `sasl`

Sorry! This field is currently undocumented.

### `target_version`

Sorry! This field is currently undocumented.

### `timeout`

Sorry! This field is currently undocumented.

### `tls`

Sorry! This field is currently undocumented.

### `topic`

Sorry! This field is currently undocumented.

