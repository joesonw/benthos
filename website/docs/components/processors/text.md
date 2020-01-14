---
title: text
type: processor
---

```yaml
text:
  arg: ""
  operator: trim_space
  parts: []
  value: ""
```

Performs text based mutations on payloads.

This processor will interpolate functions within the `value` field,
you can find a list of functions [here](../config_interpolation.md#functions).

Value interpolations are resolved once per message batch, in order to resolve it
for each message of the batch place it within a
[`for_each`](#for_each) processor:

``` yaml
for_each:
- text:
    operator: set
    value: ${!json_field:document.content}
```

### Operators

#### `append`

Appends text to the end of the payload.

#### `escape_url_query`

Escapes text so that it is safe to place within the query section of a URL.

#### `unescape_url_query`

Unescapes text that has been url escaped.

#### `find_regexp`

Extract the matching section of the argument regular expression in a message.

#### `prepend`

Prepends text to the beginning of the payload.

#### `quote`

Returns a doubled-quoted string, using escape sequences (\t, \n, \xFF, \u0100)
for control characters and other non-printable characters.

#### `regexp_expand`

Expands each matched occurrence of the argument regular expression according to
a template specified with the `value` field, and replaces the message
with the aggregated results.

Inside the template $ signs are interpreted as submatch expansions, e.g. $1
represents the text of the first submatch.

For example, given the following config:

```yaml
  - text:
      operator: regexp_expand
      arg: "(?m)(?P<key>\\w+):\\s+(?P<value>\\w+)$"
      value: "$key=$value\n"
```

And a message containing:

```text
option1: value1
# comment line
option2: value2
```

The resulting payload would be:

```text
option1=value1
option2=value2
```

#### `replace`

Replaces all occurrences of the argument in a message with a value.

#### `replace_regexp`

Replaces all occurrences of the argument regular expression in a message with a
value. Inside the value $ signs are interpreted as submatch expansions, e.g. $1
represents the text of the first submatch.

#### `set`

Replace the contents of a message entirely with a value.

#### `strip_html`

Removes all HTML tags from a message.

#### `to_lower`

Converts all text into lower case.

#### `to_upper`

Converts all text into upper case.

#### `trim`

Removes all leading and trailing occurrences of characters within the arg field.

#### `trim_space`

Removes all leading and trailing whitespace from the payload.

#### `unquote`

Unquotes a single, double, or back-quoted string literal

