---
title: hdfs
type: output
---

```yaml
hdfs:
  directory: ""
  hosts:
  - localhost:9000
  max_in_flight: 1
  path: ${!count:files}-${!timestamp_unix_nano}.txt
  user: benthos_hdfs
```

Sends message parts as files to a HDFS directory. Each file is written
with the path specified with the 'path' field, in order to have a different path
for each object you should use function interpolations described
[here](../config_interpolation.md#functions). When sending batched messages the
interpolations are performed per message part.

## Fields

### `directory`

Sorry! This field is currently undocumented.

### `hosts`

Sorry! This field is currently undocumented.

### `max_in_flight`

Sorry! This field is currently undocumented.

### `path`

Sorry! This field is currently undocumented.

### `user`

Sorry! This field is currently undocumented.

