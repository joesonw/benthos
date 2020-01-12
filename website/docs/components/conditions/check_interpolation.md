---
title: check_interpolation
type: condition
---

```yaml
check_interpolation:
  condition: {}
  value: ""
```

Resolves a string containing
[function interpolations](../config_interpolation.md#functions) and then tests
the result against a child condition.

For example, you could use this to test against the size of a message batch:

``` yaml
check_interpolation:
  value: ${!batch_size}
  condition:
    number:
      operator: greater_than
      arg: 1
```

## Fields

### `condition`

Sorry! This field is currently undocumented.

### `value`

Sorry! This field is currently undocumented.

