---
title: metadata
type: condition
---

```yaml
metadata:
  arg: ""
  key: ""
  operator: equals_cs
  part: 0
```

Metadata is a condition that checks metadata keys of a message part against an
operator from the following list:

### `enum`

Checks whether the contents of a metadata key matches one of the defined enum
values.

```yaml
metadata:
  operator: enum
  part: 0
  key: foo
  arg:
    - bar
    - baz
    - qux
    - quux
```

### `equals`

Checks whether the contents of a metadata key matches an argument. This operator
is case insensitive.

```yaml
metadata:
  operator: equals
  part: 0
  key: foo
  arg: bar
```

### `equals_cs`

Checks whether the contents of a metadata key matches an argument. This operator
is case sensitive.

```yaml
metadata:
  operator: equals_cs
  part: 0
  key: foo
  arg: BAR
```

### `exists`

Checks whether a metadata key exists.

```yaml
metadata:
  operator: exists
  part: 0
  key: foo
```

### `greater_than`

Checks whether the contents of a metadata key, parsed as a floating point
number, is greater than an argument. Returns false if the metadata value cannot
be parsed into a number.

```yaml
metadata:
  operator: greater_than
  part: 0
  key: foo
  arg: 3
```

### `has_prefix`

Checks whether the contents of a metadata key match one of the provided prefixes.
The arg field can either be a singular prefix string or a list of prefixes.

```yaml
metadata:
  operator: has_prefix
  part: 0
  key: foo
  arg:
    - foo
    - bar
    - baz
```

### `less_than`

Checks whether the contents of a metadata key, parsed as a floating point
number, is less than an argument. Returns false if the metadata value cannot be
parsed into a number.

```yaml
metadata:
  operator: less_than
  part: 0
  key: foo
  arg: 3
```

### `regexp_partial`

Checks whether any section of the contents of a metadata key matches a regular
expression (RE2 syntax).

```yaml
metadata:
  operator: regexp_partial
  part: 0
  key: foo
  arg: "1[a-z]2"
```

### `regexp_exact`

Checks whether the contents of a metadata key exactly matches a regular expression 
(RE2 syntax).

```yaml
metadata:
  operator: regexp_partial
  part: 0
  key: foo
  arg: "1[a-z]2"
```


