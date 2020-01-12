---
title: nats_stream
type: input
---

```yaml
nats_stream: {}
```

Subscribe to a NATS Stream subject, which is at-least-once. Joining a queue is
optional and allows multiple clients of a subject to consume using queue
semantics.

Tracking and persisting offsets through a durable name is also optional and
works with or without a queue. If a durable name is not provided then subjects
are consumed from the most recently published message.

When a consumer closes its connection it unsubscribes, when all consumers of a
durable queue do this the offsets are deleted. In order to avoid this you can
stop the consumers from unsubscribing by setting the field
`unsubscribe_on_close` to `false`.

### Metadata

This input adds the following metadata fields to each message:

``` text
- nats_stream_subject
- nats_stream_sequence
```

You can access these metadata fields using
[function interpolation](../config_interpolation.md#metadata).

