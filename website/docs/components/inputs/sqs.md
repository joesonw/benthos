---
title: sqs
type: input
---

```yaml
sqs:
  credentials:
    id: ""
    profile: ""
    role: ""
    role_external_id: ""
    secret: ""
    token: ""
  delete_message: true
  endpoint: ""
  max_number_of_messages: 1
  region: eu-west-1
  timeout: 5s
  url: ""
```

Receive messages from an Amazon SQS URL, only the body is extracted into
messages.

Messages consumed by this input can be processed in parallel, meaning a single
instance of this input can utilise any number of threads within a
`pipeline` section of a config.

### Credentials

By default Benthos will use a shared credentials file when connecting to AWS
services. It's also possible to set them explicitly at the component level,
allowing you to transfer data across accounts. You can find out more
[in this document](../aws.md).

### Metadata

This input adds the following metadata fields to each message:

```text
- sqs_message_id
- sqs_receipt_handle
- sqs_approximate_receive_count
- All message attributes
```

You can access these metadata fields using
[function interpolation](../config_interpolation.md#metadata).

