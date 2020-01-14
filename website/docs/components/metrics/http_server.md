---
title: http_server
type: metrics
---

```yaml
http_server:
  prefix: benthos
```

It is possible to expose metrics without an aggregator service by having Benthos
serve them as a JSON object at the endpoints `/stats` and `/metrics`.
This is useful for quickly debugging a pipeline.

The object takes the form of a hierarchical representation of the dot paths for
each metric combined. So, for example, if Benthos exposed two metric counters
`foo.bar` and `bar.baz` then the resulting object might look like
this:

``` json
{
	"foo": {
		"bar": 9
	},
	"bar": {
		"baz": 3
	}
}
```

