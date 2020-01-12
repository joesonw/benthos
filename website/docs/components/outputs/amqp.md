---
title: amqp
type: output
---

```yaml
amqp:
  exchange: benthos-exchange
  exchange_declare:
    durable: true
    enabled: false
    type: direct
  immediate: false
  key: benthos-key
  mandatory: false
  max_in_flight: 1
  persistent: false
  tls:
    client_certs: []
    enabled: false
    root_cas_file: ""
    skip_cert_verify: false
  url: amqp://guest:guest@localhost:5672/
```

DEPRECATED: This output is deprecated and scheduled for removal in Benthos V4.
Please use [`amqp_0_9`](#amqp_0_9) instead.

## Fields

### `exchange`

Sorry! This field is currently undocumented.

### `exchange_declare`

Sorry! This field is currently undocumented.

### `immediate`

Sorry! This field is currently undocumented.

### `key`

Sorry! This field is currently undocumented.

### `mandatory`

Sorry! This field is currently undocumented.

### `max_in_flight`

Sorry! This field is currently undocumented.

### `persistent`

Sorry! This field is currently undocumented.

### `tls`

Sorry! This field is currently undocumented.

### `url`

Sorry! This field is currently undocumented.

