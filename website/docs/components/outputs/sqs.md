---
title: sqs
type: output
---

```yaml
sqs:
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
  message_deduplication_id: ""
  message_group_id: ""
  region: eu-west-1
  url: ""
```

Sends messages to an SQS queue. Metadata values are sent along with the payload
as attributes with the data type String. If the number of metadata values in a
message exceeds the message attribute limit (10) then the top ten keys ordered
alphabetically will be selected.

The fields `message_group_id` and `message_deduplication_id` can be
set dynamically using
[function interpolations](../config_interpolation.md#functions), which are
resolved individually for each message of a batch.

### Credentials

By default Benthos will use a shared credentials file when connecting to AWS
services. It's also possible to set them explicitly at the component level,
allowing you to transfer data across accounts. You can find out more
[in this document](../aws.md).

