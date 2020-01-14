---
title: group_by_value
type: processor
---

```yaml
group_by_value:
  value: ${!metadata:example}
```

Splits a batch of messages into N batches, where each resulting batch contains a
group of messages determined by a
[function interpolated string](../config_interpolation.md#functions) evaluated
per message. This allows you to group messages using arbitrary fields within
their content or metadata, process them individually, and send them to unique
locations as per their group.

For example, if we were consuming Kafka messages and needed to group them by
their key, archive the groups, and send them to S3 with the key as part of the
path we could achieve that with the following:

``` yaml
pipeline:
  processors:
  - group_by_value:
      value: ${!metadata:kafka_key}
  - archive:
      format: tar
  - compress:
      algorithm: gzip
output:
  s3:
    bucket: TODO
    path: docs/${!metadata:kafka_key}/${!count:files}-${!timestamp_unix_nano}.tar.gz
```

## Fields

### `value`

Sorry! This field is currently undocumented.

