---
title: text
type: condition
---

```yaml
text:
  arg: ""
  operator: equals_cs
  part: 0
```

Text is a condition that checks the contents of a message as plain text against
a logical operator and an argument.

It's possible to use the [`check_field`](#check_field) and
[`check_interpolation`](#check_interpolation) conditions to check a
text condition against arbitrary metadata or fields of messages. For example,
you can test a text condition against a JSON field `foo.bar` with:

``` yaml
check_field:
  path: foo.bar
  condition:
    text:
      operator: enum
      arg:
      - foo
      - bar
```

Available logical operators are:

### `equals_cs`

Checks whether the content equals the argument (case sensitive.)

### `equals`

Checks whether the content equals the argument under unicode case-folding (case
insensitive.)

### `contains_cs`

Checks whether the content contains the argument (case sensitive.)

### `contains`

Checks whether the content contains the argument under unicode case-folding
(case insensitive.)

### `is`

Checks whether the content meets the characteristic of a type specified in 
the argument field. Supported types are `ip`, `ipv4`, `ipv6`.

### `prefix_cs`

Checks whether the content begins with the argument (case sensitive.)

### `prefix`

Checks whether the content begins with the argument under unicode case-folding
(case insensitive.)

### `suffix_cs`

Checks whether the content ends with the argument (case sensitive.)

### `suffix`

Checks whether the content ends with the argument under unicode case-folding
(case insensitive.)

### `regexp_partial`

Checks whether any section of the content matches a regular expression (RE2
syntax).

### `regexp_exact`

Checks whether the content exactly matches a regular expression (RE2 syntax).

### `enum`

Checks whether the content matches any entry of a list of arguments, the field
`arg` must be an array for this operator, e.g.:

``` yaml
text:
  operator: enum
  arg:
  - foo
  - bar
```

