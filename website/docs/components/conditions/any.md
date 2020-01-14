---
title: any
type: condition
---

```yaml
any: {}
```

Any is a condition that tests a child condition against each message of a batch
individually. If any message passes the child condition then this condition also
passes.

For example, if we wanted to check that at least one message of a batch contains
the word 'foo' we could use this config:

``` yaml
any:
  text:
    operator: contains
    arg: foo
```

