---
title: kinesis
type: output
---

```yaml
kinesis:
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
  hash_key: ""
  max_in_flight: 1
  max_retries: 0
  partition_key: ""
  region: eu-west-1
  stream: ""
```

Sends messages to a Kinesis stream.

Both the `partition_key`(required) and `hash_key` (optional)
fields can be dynamically set using function interpolations described
[here](../config_interpolation.md#functions). When sending batched messages the
interpolations are performed per message part.

### Credentials

By default Benthos will use a shared credentials file when connecting to AWS
services. It's also possible to set them explicitly at the component level,
allowing you to transfer data across accounts. You can find out more
[in this document](../aws.md).

