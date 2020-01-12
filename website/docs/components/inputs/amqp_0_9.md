---
title: amqp_0_9
type: input
---

```yaml
amqp_0_9:
  bindings_declare: []
  consumer_tag: benthos-consumer
  prefetch_count: 10
  prefetch_size: 0
  queue: benthos-queue
  queue_declare:
    durable: true
    enabled: false
  tls:
    client_certs: []
    enabled: false
    root_cas_file: ""
    skip_cert_verify: false
  url: amqp://guest:guest@localhost:5672/
```

Connects to an AMQP (0.91) queue. AMQP is a messaging protocol used by various
message brokers, including RabbitMQ.

It's possible for this input type to declare the target queue by setting
`queue_declare.enabled` to `true`, if the queue already exists then
the declaration passively verifies that they match the target fields.

Similarly, it is possible to declare queue bindings by adding objects to the
`bindings_declare` array. Binding declare objects take the form of:

``` yaml
{
  "exchange": "benthos-exchange",
  "key": "benthos-key"
}
```

TLS is automatic when connecting to an `amqps` URL, but custom
settings can be enabled in the `tls` section.

### Metadata

This input adds the following metadata fields to each message:

``` text
- amqp_content_type
- amqp_content_encoding
- amqp_delivery_mode
- amqp_priority
- amqp_correlation_id
- amqp_reply_to
- amqp_expiration
- amqp_message_id
- amqp_timestamp
- amqp_type
- amqp_user_id
- amqp_app_id
- amqp_consumer_tag
- amqp_delivery_tag
- amqp_redelivered
- amqp_exchange
- amqp_routing_key
- All existing message headers, including nested headers prefixed with the key
  of their respective parent.
```

You can access these metadata fields using
[function interpolation](../config_interpolation.md#metadata).

## Fields

### `bindings_declare`

Sorry! This field is currently undocumented.

### `consumer_tag`

Sorry! This field is currently undocumented.

### `prefetch_count`

Sorry! This field is currently undocumented.

### `prefetch_size`

Sorry! This field is currently undocumented.

### `queue`

Sorry! This field is currently undocumented.

### `queue_declare`

Sorry! This field is currently undocumented.

### `tls`

Sorry! This field is currently undocumented.

### `url`

Sorry! This field is currently undocumented.

