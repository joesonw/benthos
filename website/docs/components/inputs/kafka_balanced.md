---
title: kafka_balanced
type: input
---


import Tabs from '@theme/Tabs';

<Tabs defaultValue="common" values={[
  { label: 'Common', value: 'common', },
  { label: 'Advanced', value: 'advanced', },
]}>

import TabItem from '@theme/TabItem';

<TabItem value="common">

```yaml
kafka_balanced:
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
  group:
    heartbeat_interval: 3s
    rebalance_timeout: 60s
    session_timeout: 10s
  max_processing_period: 100ms
  sasl:
    enabled: false
    password: ""
    user: ""
  start_from_oldest: true
  target_version: 1.0.0
  topics:
  - benthos_stream
```

</TabItem>
<TabItem value="advanced">

```yaml
kafka_balanced:
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
  group:
    heartbeat_interval: 3s
    rebalance_timeout: 60s
    session_timeout: 10s
  max_processing_period: 100ms
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
  topics:
  - benthos_stream
```

</TabItem>
</Tabs>

Connects to a kafka (0.9+) server. Offsets are managed within kafka as per the
consumer group (set via config), and partitions are automatically balanced
across any members of the consumer group.

Partitions consumed by this input can be processed in parallel allowing it to
utilise <= N pipeline processing threads and parallel outputs where N is the
number of partitions allocated to this consumer.

The `batching` fields allow you to configure a
[batching policy](../batching.md#batch-policy) which will be applied per
partition. Any other batching mechanism will stall with this input due its
sequential transaction model.

The field `max_processing_period` should be set above the maximum
estimated time taken to process a message.

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

### `group`

Sorry! This field is currently undocumented.

### `max_processing_period`

Sorry! This field is currently undocumented.

### `sasl`

Sorry! This field is currently undocumented.

### `start_from_oldest`

Sorry! This field is currently undocumented.

### `target_version`

Sorry! This field is currently undocumented.

### `tls`

Custom TLS settings can be used to override system defaults. This includes
providing a collection of root certificate authorities, providing a list of
client certificates to use for client verification and skipping certificate
verification.

Client certificates can either be added by file or by raw contents.

```yaml
tls:
  client_certs:
  - cert_file: ./example.pem
    key_file: ./example.key
  enabled: true

tls:
  client_certs:
  - cert: foo
    key: bar
  enabled: true
  skip_cert_verify: true

```
### `topics`

Sorry! This field is currently undocumented.

