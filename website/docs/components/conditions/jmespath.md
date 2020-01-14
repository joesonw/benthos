---
title: jmespath
type: condition
---

```yaml
jmespath:
  part: 0
  query: ""
```

Parses a message part as a JSON blob and attempts to apply a JMESPath expression
to it, expecting a boolean response. If the response is true the condition
passes, otherwise it does not. Please refer to the
[JMESPath website](http://jmespath.org/) for information and tutorials regarding
the syntax of expressions.

For example, with the following config:

``` yaml
jmespath:
  part: 0
  query: a == 'foo'
```

If the initial jmespaths of part 0 were:

``` json
{
	"a": "foo"
}
```

Then the condition would pass.

JMESPath is traditionally used for mutating JSON, in order to do this please
instead use the [`jmespath`](../processors/README.md#jmespath)
processor.

