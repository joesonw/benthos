---
title: dedupe
type: processor
---

```yaml
dedupe:
  cache: ""
  drop_on_err: true
  hash: none
  key: ""
  parts:
  - 0
```

Dedupes message batches by caching selected (and optionally hashed) messages,
dropping batches that are already cached. The hash type can be chosen from:
none or xxhash.

This processor acts across an entire batch, in order to deduplicate individual
messages within a batch use this processor with the
[`for_each`](#for_each) processor.

Optionally, the `key` field can be populated in order to hash on a
function interpolated string rather than the full contents of messages. This
allows you to deduplicate based on dynamic fields within a message, such as its
metadata, JSON fields, etc. A full list of interpolation functions can be found
[here](../config_interpolation.md#functions).

For example, the following config would deduplicate based on the concatenated
values of the metadata field `kafka_key` and the value of the JSON
path `id` within the message contents:

``` yaml
dedupe:
  cache: foocache
  key: ${!metadata:kafka_key}-${!json_field:id}
```

Caches should be configured as a resource, for more information check out the
[documentation here](../caches/README.md).

When using this processor with an output target that might fail you should
always wrap the output within a [`retry`](../outputs/README.md#retry)
block. This ensures that during outages your messages aren't reprocessed after
failures, which would result in messages being dropped.

### Delivery Guarantees

Performing deduplication on a stream using a distributed cache voids any
at-least-once guarantees that it previously had. This is because the cache will
preserve message signatures even if the message fails to leave the Benthos
pipeline, which would cause message loss in the event of an outage at the output
sink followed by a restart of the Benthos instance.

If you intend to preserve at-least-once delivery guarantees you can avoid this
problem by using a memory based cache. This is a compromise that can achieve
effective deduplication but parallel deployments of the pipeline as well as
service restarts increase the chances of duplicates passing undetected.

