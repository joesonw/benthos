---
title: broker
type: output
---

```yaml
broker:
  batching:
    byte_size: 0
    condition:
      type: static
      static: false
    count: 1
    period: ""
  copies: 1
  outputs: []
  pattern: fan_out
```

The broker output type allows you to configure multiple output targets by
listing them:

``` yaml
output:
  broker:
    pattern: fan_out
    outputs:
    - foo:
        foo_field_1: value1
    - bar:
        bar_field_1: value2
        bar_field_2: value3
    - baz:
        baz_field_1: value4
      # Processors only applied to messages sent to baz.
      processors:
      - type: baz_processor

  # Processors applied to messages sent to all brokered outputs.
  processors:
  - type: some_processor
```

The broker pattern determines the way in which messages are allocated to outputs
and can be chosen from the following:

`fan_out`

With the fan out pattern all outputs will be sent every message that passes
through Benthos in parallel.

If an output applies back pressure it will block all subsequent messages, and if
an output fails to send a message it will be retried continuously until
completion or service shut down.

`fan_out_sequential`

Similar to the fan out pattern except outputs are written to sequentially,
meaning an output is only written to once the preceding output has confirmed
receipt of the same message.

`round_robin`

With the round robin pattern each message will be assigned a single output
following their order. If an output applies back pressure it will block all
subsequent messages. If an output fails to send a message then the message will
be re-attempted with the next input, and so on.

`greedy`

The greedy pattern results in higher output throughput at the cost of
potentially disproportionate message allocations to those outputs. Each message
is sent to a single output, which is determined by allowing outputs to claim
messages as soon as they are able to process them. This results in certain
faster outputs potentially processing more messages at the cost of slower
outputs.

`try`

The try pattern attempts to send each message to only one output, starting from
the first output on the list. If an output attempt fails then the broker
attempts to send to the next output in the list and so on.

This pattern is useful for triggering events in the case where certain output
targets have broken. For example, if you had an output type `http_client`
but wished to reroute messages whenever the endpoint becomes unreachable you
could use a try broker.

### Batching

It's possible to configure a [batch policy](../batching.md#batch-policy) with a
broker using the `batching` fields, allowing you to create batches
after your processing stages. Some inputs do not support broker based batching
and specify this in their documentation.

### Utilising More Outputs

When using brokered outputs with patterns such as round robin or greedy it is
possible to have multiple messages in-flight at the same time. In order to fully
utilise this you either need to have a greater number of input sources than
output sources [or use a buffer](../buffers/README.md).

### Processors

It is possible to configure [processors](../processors/README.md) at the broker
level, where they will be applied to _all_ child outputs, as well as on the
individual child outputs. If you have processors at both the broker level _and_
on child outputs then the broker processors will be applied _before_ the child
nodes processors.

