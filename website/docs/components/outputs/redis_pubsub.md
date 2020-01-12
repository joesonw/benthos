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

