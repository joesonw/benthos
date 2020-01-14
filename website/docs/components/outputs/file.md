---
title: file
type: output
---

```yaml
file:
  delimiter: ""
  path: ""
```

The file output type simply appends all messages to an output file. Single part
messages are printed with a delimiter (defaults to '\n' if left empty).
Multipart messages are written with each part delimited, with the final part
followed by two delimiters, e.g. a multipart message [ "foo", "bar", "baz" ]
would be written as:

foo\n
bar\n
baz\n\n

## Fields

### `delimiter`

Sorry! This field is currently undocumented.

### `path`

Sorry! This field is currently undocumented.

