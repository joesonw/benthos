---
title: inproc
type: output
---

```yaml
inproc: ""
```

Sends data directly to Benthos inputs by connecting to a unique ID. This allows
you to hook up isolated streams whilst running Benthos in
[`--streams` mode](../streams/README.md) mode, it is NOT recommended
that you connect the inputs of a stream with an output of the same stream, as
feedback loops can lead to deadlocks in your message flow.

It is possible to connect multiple inputs to the same inproc ID, but only one
output can connect to an inproc ID, and will replace existing outputs if a
collision occurs.

