---
title: jaeger
type: tracer
---

```yaml
jaeger:
  agent_address: localhost:6831
  flush_interval: ""
  sampler_manager_address: ""
  sampler_param: 1
  sampler_type: const
  service_name: benthos
  tags: {}
```

Send spans to a [Jaeger](https://www.jaegertracing.io/) agent.

Available sampler types are: const, probabilistic, ratelimiting and remote.

