---
title: group_by
type: processor
---

```yaml
group_by: []
```

Splits a batch of messages into N batches, where each resulting batch contains a
group of messages determined by conditions that are applied per message of the
original batch. Once the groups are established a list of processors are applied
to their respective grouped batch, which can be used to label the batch as per
their grouping.

Each group is configured in a list with a condition and a list of processors:

``` yaml
group_by:
- condition:
    static: true
  processors:
  - type: noop
```

Messages are added to the first group that passes and can only belong to a
single group. Messages that do not pass the conditions of any group are placed
in a final batch with no processors applied.

For example, imagine we have a batch of messages that we wish to split into two
groups - the foos and the bars - which should be sent to different output
destinations based on those groupings. We also need to send the foos as a tar
gzip archive. For this purpose we can use the `group_by` processor
with a [`switch`](../outputs/README.md#switch) output:

``` yaml
pipeline:
  processors:
  - group_by:
    - condition:
        text:
          operator: contains
          arg: "this is a foo"
      processors:
      - archive:
          format: tar
      - compress:
          algorithm: gzip
      - metadata:
          operator: set
          key: grouping
          value: foo
output:
  switch:
    outputs:
    - output:
        type: foo_output
      condition:
        metadata:
          operator: equals
          key: grouping
          arg: foo
    - output:
        type: bar_output
```

Since any message that isn't a foo is a bar, and bars do not require their own
processing steps, we only need a single grouping configuration.

