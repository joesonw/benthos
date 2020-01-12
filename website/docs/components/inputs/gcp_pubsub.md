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
[function interpolation](../config_interpolation.md#metadata).

## Fields

### `max_outstanding_bytes`

Sorry! This field is currently undocumented.

### `max_outstanding_messages`

Sorry! This field is currently undocumented.

### `project`

Sorry! This field is currently undocumented.

### `subscription`

Sorry! This field is currently undocumented.

