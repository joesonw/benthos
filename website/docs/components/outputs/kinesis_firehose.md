---
title: kinesis_firehose
type: output
---

```yaml
kinesis_firehose:
  backoff:
    initial_interval: 1s
    max_elapsed_time: 30s
    max_interval: 5s
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
  endpoint: ""
  max_in_flight: 1
  max_retries: 0
  region: eu-west-1
  stream: ""
```

Sends messages to a Kinesis Firehose delivery stream.

### Credentials

By default Benthos will use a shared credentials file when connecting to AWS
services. It's also possible to set them explicitly at the component level,
allowing you to transfer data across accounts. You can find out more
[in this document](../aws.md).

## Fields

### `backoff`

Sorry! This field is currently undocumented.

### `batching`

Sorry! This field is currently undocumented.

### `credentials`

Sorry! This field is currently undocumented.

### `endpoint`

Sorry! This field is currently undocumented.

### `max_in_flight`

Sorry! This field is currently undocumented.

### `max_retries`

Sorry! This field is currently undocumented.

### `region`

Sorry! This field is currently undocumented.

### `stream`

Sorry! This field is currently undocumented.

