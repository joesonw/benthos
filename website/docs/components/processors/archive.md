---
title: archive
type: processor
---

```yaml
archive:
  format: binary
  path: ${!count:files}-${!timestamp_unix_nano}.txt
```

Archives all the messages of a batch into a single message according to the
selected archive format. Supported archive formats are:
`tar`, `zip`, `binary`, `lines` and `json_array`.

Some archive formats (such as tar, zip) treat each archive item (message part)
as a file with a path. Since message parts only contain raw data a unique path
must be generated for each part. This can be done by using function
interpolations on the 'path' field as described
[here](../config_interpolation.md#functions). For types that aren't file based
(such as binary) the file field is ignored.

The `json_array` format attempts to JSON parse each message and append
the result to an array, which becomes the contents of the resulting message.

The resulting archived message adopts the metadata of the _first_ message part
of the batch.

## Fields

### `format`

Sorry! This field is currently undocumented.

### `path`

Sorry! This field is currently undocumented.

