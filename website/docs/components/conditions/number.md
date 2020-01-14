---
title: number
type: condition
---

```yaml
number:
  arg: 0
  operator: equals
  part: 0
```

Number is a condition that checks the contents of a message parsed as a 64-bit
floating point number against a logical operator and an argument.

It's possible to use the [`check_field`](#check_field) and
[`check_interpolation`](#check_interpolation) conditions to check a
number condition against arbitrary metadata or fields of messages. For example,
you can test a number condition against the size of a message batch with:

``` yaml
check_interpolation:
  value: ${!batch_size}
  condition:
    number:
      operator: greater_than
      arg: 1
```

Available logical operators are:

### `equals`

Checks whether the value equals the argument.

### `greater_than`

Checks whether the value is greater than the argument. Returns false if the
value cannot be parsed as a number.

### `less_than`

Checks whether the value is less than the argument. Returns false if the value
cannot be parsed as a number.

