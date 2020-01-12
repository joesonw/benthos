---
title: dynamodb
type: output
---

```yaml
dynamodb:
  backoff:
    initial_interval: 1s
    max_elapsed_time: 30s
    max_interval: 5s
  batching:
    byte_size: 0
    condition:
      type: static
      static: false
    count: 1
    period: ""
  credentials:
    id: ""
    profile: ""
    role: ""
    role_external_id: ""
    secret: ""
    token: ""
  endpoint: ""
  json_map_columns: {}
  max_in_flight: 1
  max_retries: 3
  region: eu-west-1
  string_columns: {}
  table: ""
  ttl: ""
  ttl_key: ""
```

Inserts items into a DynamoDB table.

The field `string_columns` is a map of column names to string values,
where the values are
[function interpolated](../config_interpolation.md#functions) per message of a
batch. This allows you to populate string columns of an item by extracting
fields within the document payload or metadata like follows:

``` yaml
string_columns:
  id: ${!json_field:id}
  title: ${!json_field:body.title}
  topic: ${!metadata:kafka_topic}
  full_content: ${!content}
```

The field `json_map_columns` is a map of column names to json paths,
where the [dot path](../field_paths.md) is extracted from each document and
converted into a map value. Both an empty path and the path `.` are
interpreted as the root of the document. This allows you to populate map columns
of an item like follows:

``` yaml
json_map_columns:
  user: path.to.user
  whole_document: .
```

A column name can be empty:

``` yaml
json_map_columns:
  "": .
```

In which case the top level document fields will be written at the root of the
item, potentially overwriting previously defined column values. If a path is not
found within a document the column will not be populated.

### Credentials

By default Benthos will use a shared credentials file when connecting to AWS
services. It's also possible to set them explicitly at the component level,
allowing you to transfer data across accounts. You can find out more
[in this document](../aws.md).

## Fields

### `backoff`

Sorry! This field is currently undocumented.

### `batching`

Sorry! This field is currently undocumented.

### `credentials`

Sorry! This field is currently undocumented.

### `endpoint`

Sorry! This field is currently undocumented.

### `json_map_columns`

Sorry! This field is currently undocumented.

### `max_in_flight`

Sorry! This field is currently undocumented.

### `max_retries`

Sorry! This field is currently undocumented.

### `region`

Sorry! This field is currently undocumented.

### `string_columns`

Sorry! This field is currently undocumented.

### `table`

Sorry! This field is currently undocumented.

### `ttl`

Sorry! This field is currently undocumented.

### `ttl_key`

Sorry! This field is currently undocumented.

