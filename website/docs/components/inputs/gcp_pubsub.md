---
title: gcp_pubsub
type: input
---

```yaml
gcp_pubsub:
  max_outstanding_bytes: 1000000000
  max_outstanding_messages: 1000
  project: ""
  subscription: ""
```

Consumes messages from a GCP Cloud Pub/Sub subscription.

### Metadata

This input adds the following metadata fields to each message:

``` text
- gcp_pubsub_publish_time_unix
- All message attributes
```

You can access these metadata fields using
[function interpolation](/docs/configuration/interpolation#metadata).

## Fields

### `project`

`string` The project ID of the target subscription.
### `subscription`

`string` The target subscription ID.
### `max_outstanding_messages`

`number` The maximum number of outstanding pending messages to be consumed at a given time.
### `max_outstanding_bytes`

`number` The maximum number of outstanding pending messages to be consumed measured in bytes.
