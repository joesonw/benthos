---
title: nats
type: input
---

```yaml
nats:
  prefetch_count: 32
  queue: benthos_queue
  subject: benthos_messages
  urls:
  - nats://127.0.0.1:4222
```

Subscribe to a NATS subject. NATS is at-most-once, if you need at-least-once
behaviour then look at NATS Stream.

The urls can contain username/password semantics. e.g.
nats://derek:pass@localhost:4222

### Metadata

This input adds the following metadata fields to each message:

``` text
- nats_subject
```

You can access these metadata fields using
[function interpolation](../config_interpolation.md#metadata).

## Fields

### `prefetch_count`

Sorry! This field is currently undocumented.

### `queue`

Sorry! This field is currently undocumented.

### `subject`

Sorry! This field is currently undocumented.

### `urls`

Sorry! This field is currently undocumented.

