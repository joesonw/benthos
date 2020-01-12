---
title: nanomsg
type: input
---

```yaml
nanomsg:
  bind: true
  poll_timeout: 5s
  reply_timeout: 5s
  socket_type: PULL
  sub_filters: []
  urls:
  - tcp://*:5555
```

The scalability protocols are common communication patterns. This input should
be compatible with any implementation, but specifically targets Nanomsg.

Currently only PULL and SUB sockets are supported.

