---
title: mqtt
type: output
---

```yaml
mqtt:
  client_id: benthos_output
  max_in_flight: 1
  password: ""
  qos: 1
  topic: benthos_topic
  urls:
  - tcp://localhost:1883
  user: ""
```

Pushes messages to an MQTT broker.

The `topic` field can be dynamically set using function interpolations
described [here](../config_interpolation.md#functions). When sending batched
messages these interpolations are performed per message part.

