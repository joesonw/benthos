---
title: metric
type: processor
---

```yaml
metric:
  labels: {}
  path: ""
  type: counter
  value: ""
```

Expose custom metrics by extracting values from message batches. This processor
executes once per batch, in order to execute once per message place it within a
[`for_each`](#for_each) processor:

``` yaml
for_each:
- metric:
    type: counter_by
    path: count.custom.field
    value: ${!json_field:field.some.value}
```

The `path` field should be a dot separated path of the metric to be
set and will automatically be converted into the correct format of the
configured metric aggregator.

The `value` field can be set using function interpolations described
[here](../config_interpolation.md#functions) and is used according to the
specific type.

### Types

#### `counter`

Increments a counter by exactly 1, the contents of `value` are ignored
by this type.

#### `counter_parts`

Increments a counter by the number of parts within the message batch, the
contents of `value` are ignored by this type.

#### `counter_by`

If the contents of `value` can be parsed as a positive integer value
then the counter is incremented by this value.

For example, the following configuration will increment the value of the
`count.custom.field` metric by the contents of `field.some.value`:

``` yaml
metric:
  type: counter_by
  path: count.custom.field
  value: ${!json_field:field.some.value}
```

#### `gauge`

If the contents of `value` can be parsed as a positive integer value
then the gauge is set to this value.

For example, the following configuration will set the value of the
`gauge.custom.field` metric to the contents of `field.some.value`:

``` yaml
metric:
  type: gauge
  path: gauge.custom.field
  value: ${!json_field:field.some.value}
```

#### `timing`

Equivalent to `gauge` where instead the metric is a timing.

### Labels

Some metrics aggregators, such as Prometheus, support arbitrary labels, in which
case the `labels` field can be used in order to create them. Label
values can also be set using function interpolations in order to dynamically
populate them with context about the message.

