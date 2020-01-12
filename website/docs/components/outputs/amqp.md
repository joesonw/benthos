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

