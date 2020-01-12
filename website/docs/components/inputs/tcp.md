---
title: tcp
type: input
---

```yaml
tcp:
  address: localhost:4194
  delimiter: ""
  max_buffer: 1e+06
  multipart: false
```

Connects to a TCP server and consumes a continuous stream of messages.

If multipart is set to false each line of data is read as a separate message. If
multipart is set to true each line is read as a message part, and an empty line
indicates the end of a message.

Messages consumed by this input can be processed in parallel, meaning a single
instance of this input can utilise any number of threads within a
`pipeline` section of a config.

If the delimiter field is left empty then line feed (\n) is used.

## Fields

### `address`

Sorry! This field is currently undocumented.

### `delimiter`

Sorry! This field is currently undocumented.

### `max_buffer`

Sorry! This field is currently undocumented.

### `multipart`

Sorry! This field is currently undocumented.

