---
title: nsq
type: output
---

```yaml
nsq:
  max_in_flight: 1
  nsqd_tcp_address: localhost:4150
  topic: benthos_messages
  user_agent: benthos_producer
```

Publish to an NSQ topic. The `topic` field can be dynamically set
using function interpolations described
[here](../config_interpolation.md#functions). When sending batched messages
these interpolations are performed per message part.

## Fields

### `max_in_flight`

Sorry! This field is currently undocumented.

### `nsqd_tcp_address`

Sorry! This field is currently undocumented.

### `topic`

Sorry! This field is currently undocumented.

### `user_agent`

Sorry! This field is currently undocumented.

