---
title: unarchive
type: processor
---

```yaml
unarchive:
  format: binary
  parts: []
```

Unarchives messages according to the selected archive format into multiple
messages within a batch. Supported archive formats are:
`tar`, `zip`, `binary`, `lines`, `json_documents`, `json_array` and `json_map`.

When a message is unarchived the new messages replaces the original message in
the batch. Messages that are selected but fail to unarchive (invalid format)
will remain unchanged in the message batch but will be flagged as having failed.

The `json_documents` format attempts to parse the message as a stream
of concatenated JSON documents. Each parsed document is expanded into a new
message.

The `json_array` format attempts to parse the message as a JSON array
and for each element of the array expands its contents into a new message.

The `json_map` format attempts to parse the message as a JSON map
and for each element of the map expands its contents into a new message.

For the unarchive formats that contain file information (tar, zip), a metadata
field is added to each message called `archive_filename` with the
extracted filename.

For the `json_map` format, a metadata field is added to each message
called `archive_key` with the relevant key from the top-level map.

## Fields

### `format`

Sorry! This field is currently undocumented.

### `parts`

Sorry! This field is currently undocumented.

