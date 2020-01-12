---
title: files
type: input
---

```yaml
files:
  path: ""
```

Reads files from a path, where each discrete file will be consumed as a single
message payload. The path can either point to a single file (resulting in only a
single message) or a directory, in which case the directory will be walked and
each file found will become a message.

### Metadata

This input adds the following metadata fields to each message:

``` text
- path
```

You can access these metadata fields using
[function interpolation](../config_interpolation.md#metadata).

## Fields

### `path`

Sorry! This field is currently undocumented.

