---
title: hdfs
type: input
---

```yaml
hdfs:
  directory: ""
  hosts:
  - localhost:9000
  user: benthos_hdfs
```

Reads files from a HDFS directory, where each discrete file will be consumed as
a single message payload.

### Metadata

This input adds the following metadata fields to each message:

``` text
- hdfs_name
- hdfs_path
```

You can access these metadata fields using
[function interpolation](../config_interpolation.md#metadata).

## Fields

### `directory`

Sorry! This field is currently undocumented.

### `hosts`

Sorry! This field is currently undocumented.

### `user`

Sorry! This field is currently undocumented.

