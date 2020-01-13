---
title: broker
type: input
---

```yaml
broker:
  copies: 1
  inputs: []
  batching:
    byte_size: 0
    condition:
      type: static
      static: false
    count: 1
    period: ""
```

The broker type allows you to combine multiple inputs, where each input will be
read in parallel. A broker type is configured with its own list of input
configurations and a field to specify how many copies of the list of inputs
should be created.

Adding more input types allows you to merge streams from multiple sources into
one. For example, reading from both RabbitMQ and Kafka:

``` yaml
input:
  broker:
    copies: 1
    inputs:
    - amqp:
        url: amqp://guest:guest@localhost:5672/
        consumer_tag: benthos-consumer
        queue: benthos-queue

      # Optional list of input specific processing steps
      processors:
        - jmespath:
            query: '{ message: @, meta: { link_count: length(links) } }'

    - kafka:
        addresses:
        - localhost:9092
        client_id: benthos_kafka_input
        consumer_group: benthos_consumer_group
        partition: 0
        topic: benthos_stream
```

If the number of copies is greater than zero the list will be copied that number
of times. For example, if your inputs were of type foo and bar, with 'copies'
set to '2', you would end up with two 'foo' inputs and two 'bar' inputs.

### Batching

It's possible to configure a [batch policy](/docs/configuration/batching#batch-policy)
with a broker using the `batching` fields. When doing this the feeds
from all child inputs are combined. Some inputs do not support broker based
batching and specify this in their documentation.

### Processors

It is possible to configure [processors](/docs/components/processors/about) at
the broker level, where they will be applied to _all_ child inputs, as well as
on the individual child inputs. If you have processors at both the broker level
_and_ on child inputs then the broker processors will be applied _after_ the
child nodes processors.

## Fields

### `copies`

`number` Whatever is specified within `inputs` will be created this many times.
### `inputs`

`array` A list of inputs to create.
### `batching`

`object` Allows you to configure a [batching policy](/docs/configuration/batching).
```yaml
# Examples

batching:
  byte_size: 5000
  period: 1s

batching:
  count: 10
  period: 1s

batching:
  condition:
    text:
      arg: END BATCH
      operator: contains
  period: 1m

```
