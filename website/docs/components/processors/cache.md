---
title: cache
type: processor
---

```yaml
cache:
  cache: ""
  key: ""
  operator: set
  parts: []
  value: ""
```

Performs operations against a [cache resource](../caches) for each message of a
batch, allowing you to store or retrieve data within message payloads.

This processor will interpolate functions within the `key` and `value`
fields individually for each message of the batch. This allows you to specify
dynamic keys and values based on the contents of the message payloads and
metadata. You can find a list of functions
[here](../config_interpolation.md#functions).

### Operators

#### `set`

Set a key in the cache to a value. If the key already exists the contents are
overridden.

#### `add`

Set a key in the cache to a value. If the key already exists the action fails
with a 'key already exists' error, which can be detected with
[processor error handling](../error_handling.md).

#### `get`

Retrieve the contents of a cached key and replace the original message payload
with the result. If the key does not exist the action fails with an error, which
can be detected with [processor error handling](../error_handling.md).

#### `delete`

Delete a key and its contents from the cache.  If the key does not exist the
action is a no-op and will not fail with an error.

### Examples

The `cache` processor can be used in combination with other processors
in order to solve a variety of data stream problems.

#### Deduplication

Deduplication can be done using the add operator with a key extracted from the
message payload, since it fails when a key already exists we can remove the
duplicates using a
[`processor_failed`](../conditions/README.md#processor_failed)
condition:

``` yaml
- cache:
    cache: TODO
    operator: add
    key: "${!json_field:message.id}"
    value: "storeme"
- filter_parts:
    type: processor_failed
```

#### Hydration

It's possible to enrich payloads with content previously stored in a cache by
using the [`process_map`](#process_map) processor:

``` yaml
- process_map:
    processors:
    - cache:
        cache: TODO
        operator: get
        key: "${!json_field:message.document_id}"
    postmap:
      message.document: .
```

