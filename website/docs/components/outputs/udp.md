---
title: udp
type: output
---

```yaml
udp:
  address: localhost:4194
```

Sends messages as a continuous stream of line delimited data over UDP by
connecting to a server.

If batched messages are sent the final message of the batch will be followed by
two line breaks in order to indicate the end of the batch.

