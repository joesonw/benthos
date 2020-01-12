---
title: redis_streams
type: output
---

```yaml
redis_streams:
  body_key: body
  max_in_flight: 1
  max_length: 0
  stream: benthos_stream
  url: tcp://localhost:6379
```

Pushes messages to a Redis (v5.0+) Stream (which is created if it doesn't
already exist) using the XADD command. It's possible to specify a maximum length
of the target stream by setting it to a value greater than 0, in which case this
cap is applied only when Redis is able to remove a whole macro node, for
efficiency.

Redis stream entries are key/value pairs, as such it is necessary to specify the
key to be set to the body of the message. All metadata fields of the message
will also be set as key/value pairs, if there is a key collision between
a metadata item and the body then the body takes precedence.

## Fields

### `body_key`

Sorry! This field is currently undocumented.

### `max_in_flight`

Sorry! This field is currently undocumented.

### `max_length`

Sorry! This field is currently undocumented.

### `stream`

Sorry! This field is currently undocumented.

### `url`

Sorry! This field is currently undocumented.

