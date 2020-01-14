---
title: process_field
type: processor
---

```yaml
process_field:
  codec: json
  parts: []
  path: ""
  processors: []
  result_type: string
```

A processor that extracts the value of a field [dot path](../field_paths.md)
within payloads according to a specified codec, applies a list of processors to
the extracted value and finally sets the field within the original payloads to
the processed result.

### Codecs

#### `json` (default)

Parses the payload as a JSON document, extracts and sets the field using a dot
notation path.

The result, according to the config field `result_type`, can be
marshalled into any of the following types:
`string` (default), `int`, `float`, `bool`, `object` (including null),
 `array` and `discard`. The discard type is a special case that
discards the result of the processing steps entirely.

It's therefore possible to use this codec without any child processors as a way
of casting string values into other types. For example, with an input JSON
document `{"foo":"10"}` it's possible to cast the value of the
field foo to an integer type with:

```yaml
process_field:
  path: foo
  result_type: int
```

#### `metadata`

Extracts and sets a metadata value identified by the path field. If the field
`result_type` is set to `discard` then the result of the processing stages
is discarded and the original metadata value is left unchanged.

### Usage

For example, with an input JSON document `{"foo":"hello world"}`
it's possible to uppercase the value of the field 'foo' by using the JSON codec
and a [`text`](#text) child processor:

```yaml
process_field:
  codec: json
  path: foo
  processors:
  - text:
      operator: to_upper
```

If the number of messages resulting from the processing steps does not match the
original count then this processor fails and the messages continue unchanged.
Therefore, you should avoid using batch and filter type processors in this list.

