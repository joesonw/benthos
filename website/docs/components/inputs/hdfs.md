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
[function interpolation](/docs/configuration/interpolation#metadata).

## Fields

### `hosts`

`array` A list of target host addresses to connect to.
### `user`

`string` A user ID to connect as.
### `directory`

`string` The directory to consume from.
