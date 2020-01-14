---
title: stdin
type: input
---

```yaml
stdin:
  delimiter: ""
  max_buffer: 1e+06
  multipart: false
```

The stdin input simply reads any data piped to stdin as messages. By default the
messages are assumed single part and are line delimited. If the multipart option
is set to true then lines are interpretted as message parts, and an empty line
indicates the end of the message.

Messages consumed by this input can be processed in parallel, meaning a single
instance of this input can utilise any number of threads within a
`pipeline` section of a config.

If the delimiter field is left empty then line feed (\n) is used.

