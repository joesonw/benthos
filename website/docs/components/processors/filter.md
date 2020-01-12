---
title: filter
type: processor
---

```yaml
filter:
  type: text
  text:
    arg: ""
    operator: equals_cs
    part: 0
```

Tests each message batch against a condition, if the condition fails then the
batch is dropped. You can find a [full list of conditions here](../conditions).

In order to filter individual messages of a batch use the
[`filter_parts`](#filter_parts) processor.

