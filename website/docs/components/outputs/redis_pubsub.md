---
title: redis_pubsub
type: output
---

```yaml
redis_pubsub:
  channel: benthos_chan
  max_in_flight: 1
  url: tcp://localhost:6379
```

Publishes messages through the Redis PubSub model. It is not possible to
guarantee that messages have been received.

This output will interpolate functions within the channel field, you
can find a list of functions [here](../config_interpolation.md#functions).

## Fields

### `channel`

Sorry! This field is currently undocumented.

### `max_in_flight`

Sorry! This field is currently undocumented.

### `url`

Sorry! This field is currently undocumented.

