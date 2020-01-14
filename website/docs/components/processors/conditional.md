---
title: conditional
type: processor
---

```yaml
conditional:
  condition:
    type: text
    text:
      arg: ""
      operator: equals_cs
      part: 0
  else_processors: []
  processors: []
```

Conditional is a processor that has a list of child `processors`,
`else_processors`, and a condition. For each message batch, if the
condition passes, the child `processors` will be applied, otherwise
the `else_processors` are applied. This processor is useful for
applying processors based on the content of message batches.

In order to conditionally process each message of a batch individually use this
processor with the [`for_each`](#for_each) processor.

You can find a [full list of conditions here](../conditions).

