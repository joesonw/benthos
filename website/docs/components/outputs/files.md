---
title: files
type: output
---

```yaml
files:
  path: ${!count:files}-${!timestamp_unix_nano}.txt
```

Writes each individual part of each message to a new file.

Message parts only contain raw data, and therefore in order to create a unique
file for each part you need to generate unique file names. This can be done by
using function interpolations on the `path` field as described
[here](../config_interpolation.md#functions). When sending batched messages
these interpolations are performed per message part.

