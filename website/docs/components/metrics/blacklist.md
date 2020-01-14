---
title: blacklist
type: metrics
---

```yaml
blacklist:
  child: {}
  paths: []
  patterns: []
```

Blacklist metric paths within Benthos so that they are not aggregated by a child
metric target.

Blacklists can either be path prefixes or regular expression patterns, if either
a path prefix or regular expression matches a metric path it will be excluded.

Metrics must be matched using dot notation even if the chosen output uses a
different form. For example, the path would be 'foo.bar' rather than 'foo_bar'
even when sending metrics to Prometheus.

### Paths

An entry in the `paths` field will check using prefix matching. This
can be used, for example, to allow none of the child specific metrics paths from
an output broker with the path `output.broker`.

### Patterns

An entry in the `patterns` field will be parsed as an RE2 regular
expression and tested against each metric path. This can be used, for example,
to allow none of the latency based metrics with the pattern
`.*\.latency`.

### Debugging

In order to see logs breaking down which metrics are registered and whether they
are blocked by your blacklists enable logging at the TRACE level.

## Fields

### `child`

Sorry! This field is currently undocumented.

### `paths`

Sorry! This field is currently undocumented.

### `patterns`

Sorry! This field is currently undocumented.

