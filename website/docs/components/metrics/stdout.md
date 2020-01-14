---
title: stdout
type: metrics
---

```yaml
stdout:
  flush_metrics: false
  push_interval: ""
  static_fields:
    '@service': benthos
```

EXPERIMENTAL: This component is considered experimental and is therefore subject
to change outside of major version releases.

It is possible to expose metrics without an aggregator service while running in
serverless mode by having Benthos output metrics as JSON objects to stdout.  This
is useful if you do not have Prometheus or Statsd endpoints and you cannot query
the Benthos API for statistics (due to the short lived nature of serverless
invocations).

A series of JSON objects are emitted (one per line) grouped by the
input/processor/output instance.  Separation into individual JSON objects instead
of a single monolithic object allows for easy ingestion into document stores such
as Elasticsearch.

If defined, metrics are pushed at the configured push_interval, otherwise they
are emitted when Benthos closes.

flush_metrics dictates whether counter and timing metrics are reset to 0 after
they are pushed out.


## Fields

### `flush_metrics`

Sorry! This field is currently undocumented.

### `push_interval`

Sorry! This field is currently undocumented.

### `static_fields`

Sorry! This field is currently undocumented.

