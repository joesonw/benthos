---
title: redis_hash
type: output
---

```yaml
redis_hash:
  fields: {}
  key: ""
  max_in_flight: 1
  url: tcp://localhost:6379
  walk_json_object: false
  walk_metadata: false
```

Sets Redis hash objects using the HMSET command.

The field `key` supports
[interpolation functions](../config_interpolation.md#functions) evaluated per
message of a batch, allowing you to create a unique key for each message.

The field `fields` allows you to specify an explicit map of field
names to interpolated values, also evaluated per message of a batch:

```yaml
redis_hash:
  url: tcp://localhost:6379
  key: ${!json_field:id}
  fields:
    topic: ${!metadata:kafka_topic}
    partition: ${!metadata:kafka_partition}
    content: ${!json_field:document.text}
```

If the field `walk_metadata` is set to `true` then Benthos
will walk all metadata fields of messages and add them to the list of hash
fields to set.

If the field `walk_json_object` is set to `true` then
Benthos will walk each message as a JSON object, extracting keys and the string
representation of their value and adds them to the list of hash fields to set.

The order of hash field extraction is as follows:

1. Metadata (if enabled)
2. JSON object (if enabled)
3. Explicit fields

Where latter stages will overwrite matching field names of a former stage.

## Fields

### `fields`

Sorry! This field is currently undocumented.

### `key`

Sorry! This field is currently undocumented.

### `max_in_flight`

Sorry! This field is currently undocumented.

### `url`

Sorry! This field is currently undocumented.

### `walk_json_object`

Sorry! This field is currently undocumented.

### `walk_metadata`

Sorry! This field is currently undocumented.

