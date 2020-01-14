---
title: statsd
type: metrics
---

```yaml
statsd:
  address: localhost:4040
  flush_period: 100ms
  network: udp
  prefix: benthos
  tag_format: legacy
```

Pushes metrics using the [StatsD protocol](https://github.com/statsd/statsd).
Supported tagging formats are 'legacy', 'none', 'datadog' and 'influxdb'.

The underlying client library has recently been updated in order to support
tagging. The tag format 'legacy' is default and causes Benthos to continue using
the old library in order to preserve backwards compatibility.

The legacy library aggregated timing metrics, so dashboards and alerts may need
to be updated when migrating to the new library.

The 'network' field is deprecated and scheduled for removal. If you currently
rely on sending Statsd metrics over TCP and want it to be supported long term
please [raise an issue](https://github.com/Jeffail/benthos/issues).

