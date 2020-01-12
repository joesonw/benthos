---
title: nanomsg
type: output
---

```yaml
nanomsg:
  bind: false
  max_in_flight: 1
  poll_timeout: 5s
  socket_type: PUSH
  urls:
  - tcp://localhost:5556
```

The scalability protocols are common communication patterns. This output should
be compatible with any implementation, but specifically targets Nanomsg.

Currently only PUSH and PUB sockets are supported.

