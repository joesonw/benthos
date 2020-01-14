---
title: rename
type: metrics
---

```yaml
rename:
  by_regexp: []
  child: {}
```

Rename metric paths as they are registered.

Metrics must be matched using dot notation even if the chosen output uses a
different form. For example, the path would be 'foo.bar' rather than 'foo_bar'
even when sending metrics to Prometheus.

### `by_regexp`

An array of objects of the form:

```yaml
  - pattern: "foo\\.([a-z]*)\\.([a-z]*)"
    value: "foo.$1"
    to_label:
      bar: $2
```

Where each pattern will be parsed as an RE2 regular expression, these
expressions are tested against each metric path, where all occurrences will be
replaced with the specified value. Inside the value $ signs are interpreted as
submatch expansions, e.g. $1 represents the first submatch.

The field `to_label` may contain any number of key/value pairs to be
added to a metric as labels, where the value may contain submatches from the
provided pattern. This allows you to extract (left-most) matched segments of the
renamed path into the label values.

For example, in order to replace the paths 'foo.bar.0.zap' and 'foo.baz.1.zap'
with 'zip.bar' and 'zip.baz' respectively, and store the respective values '0'
and '1' under the label key 'index' we could use this config:

```yaml
metrics:
  rename:
    by_regexp:
      - pattern: "foo\\.([a-z]*)\\.([a-z]*)\\.zap"
        value: "zip.$1"
        to_label:
          index: $2
    child:
      statsd:
        prefix: foo
        address: localhost:8125
```

These labels will only be injected into metrics registered without pre-existing
labels. Therefore it's currently not possible to combine labels registered from
the [`metric` processor](../processors/README.md#metric) with labels
set via renaming.

### Debugging

In order to see logs breaking down which metrics are registered and whether they
are renamed enable logging at the TRACE level.

