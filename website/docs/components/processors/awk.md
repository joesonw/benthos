---
title: awk
type: processor
---

```yaml
awk:
  codec: text
  parts: []
  program: BEGIN { x = 0 } { print $0, x; x++ }
```

Executes an AWK program on messages by feeding contents as the input based on a
codec and replaces the contents with the result. If the result is empty (nothing
is printed by the program) then the original message contents remain unchanged.

Comes with a wide range of [custom functions](./awk_functions.md) for accessing
message metadata, json fields, printing logs, etc. These functions can be
overridden by functions within the program.

### Codecs

A codec can be specified that determines how the contents of the message are fed
into the program. This does not change the custom functions.

`none`

An empty string is fed into the program. Functions can still be used in order to
extract and mutate metadata and message contents. This is useful for when your
program only uses functions and doesn't need the full text of the message to be
parsed by the program.

`text`

The full contents of the message are fed into the program as a string, allowing
you to reference tokenised segments of the message with variables ($0, $1, etc).
Custom functions can still be used with this codec.

This is the default codec as it behaves most similar to typical usage of the awk
command line tool.

`json`

No contents are fed into the program. Instead, variables are extracted from the
message by walking the flattened JSON structure. Each value is converted into a
variable by taking its full path, e.g. the object:

``` json
{
	"foo": {
		"bar": {
			"value": 10
		},
		"created_at": "2018-12-18T11:57:32"
	}
}
```

Would result in the following variable declarations:

```
foo_bar_value = 10
foo_created_at = "2018-12-18T11:57:32"
```

Custom functions can also still be used with this codec.

