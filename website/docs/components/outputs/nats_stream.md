---
title: nats_stream
type: output
---

```yaml
nats_stream:
  client_id: benthos_client
  cluster_id: test-cluster
  max_in_flight: 1
  subject: benthos_messages
  urls:
  - nats://localhost:4222
```

Publish to a NATS Stream subject.

