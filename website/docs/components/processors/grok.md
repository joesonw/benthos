---
title: grok
type: processor
---

```yaml
grok:
  named_captures_only: true
  output_format: json
  parts: []
  pattern_definitions: {}
  patterns: []
  remove_empty_values: true
  use_default_patterns: true
```

Parses message payloads by attempting to apply a list of Grok patterns, if a
pattern returns at least one value a resulting structured object is created
according to the chosen output format and will replace the payload. Currently
only json is a valid output format.

Type hints within patterns are respected, therefore with the pattern
`%{WORD:first},%{INT:second:int}` and a payload of `foo,1`
the resulting payload would be `{"first":"foo","second":1}`.

### Performance

This processor currently uses the [Go RE2](https://golang.org/s/re2syntax)
regular expression engine, which is guaranteed to run in time linear to the size
of the input. However, this property often makes it less performant than pcre
based implementations of grok. For more information see
[https://swtch.com/~rsc/regexp/regexp1.html](https://swtch.com/~rsc/regexp/regexp1.html).

