---
title: amqp_0_9
type: output
---

```yaml
amqp_0_9:
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

Sends messages to an AMQP (0.91) exchange. AMQP is a messaging protocol used by
various message brokers, including RabbitMQ. The metadata from each message are
delivered as headers.

It's possible for this output type to create the target exchange by setting
`exchange_declare.enabled` to `true`, if the exchange already exists
then the declaration passively verifies that the settings match.

Exchange type options are: direct|fanout|topic|x-custom

TLS is automatic when connecting to an `amqps` URL, but custom
settings can be enabled in the `tls` section.

The field 'key' can be dynamically set using function interpolations described
[here](../config_interpolation.md#functions).

