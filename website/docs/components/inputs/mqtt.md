---
title: mqtt
type: input
---

```yaml
mqtt:
  clean_session: true
  client_id: benthos_input
  password: ""
  qos: 1
  topics:
  - benthos_topic
  urls:
  - tcp://localhost:1883
  user: ""
```

Subscribe to topics on MQTT brokers.

### Metadata

This input adds the following metadata fields to each message:

``` text
- mqtt_duplicate
- mqtt_qos
- mqtt_retained
- mqtt_topic
- mqtt_message_id
```

You can access these metadata fields using
[function interpolation](../config_interpolation.md#metadata).

## Fields

### `clean_session`

Sorry! This field is currently undocumented.

### `client_id`

Sorry! This field is currently undocumented.

### `password`

Sorry! This field is currently undocumented.

### `qos`

Sorry! This field is currently undocumented.

### `topics`

Sorry! This field is currently undocumented.

### `urls`

Sorry! This field is currently undocumented.

### `user`

Sorry! This field is currently undocumented.

