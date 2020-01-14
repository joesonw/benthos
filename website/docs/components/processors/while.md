---
title: while
type: processor
---

```yaml
while:
  at_least_once: false
  condition:
    type: text
    text:
      arg: ""
      operator: equals_cs
      part: 0
  max_loops: 0
  processors: []
```

While is a processor that has a condition and a list of child processors. The
child processors are executed continuously on a message batch for as long as the
child condition resolves to true.

The field `at_least_once`, if true, ensures that the child processors
are always executed at least one time (like a do .. while loop.)

The field `max_loops`, if greater than zero, caps the number of loops
for a message batch to this value.

If following a loop execution the number of messages in a batch is reduced to
zero the loop is exited regardless of the condition result. If following a loop
execution there are more than 1 message batches the condition is checked against
the first batch only.

You can find a [full list of conditions here](../conditions).

