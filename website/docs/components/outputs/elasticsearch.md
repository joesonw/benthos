---
title: elasticsearch
type: output
---

```yaml
elasticsearch:
  aws:
    credentials:
      id: ""
      profile: ""
      role: ""
      role_external_id: ""
      secret: ""
      token: ""
    enabled: false
    endpoint: ""
    region: eu-west-1
  backoff:
    initial_interval: 1s
    max_elapsed_time: 30s
    max_interval: 5s
  basic_auth:
    enabled: false
    password: ""
    username: ""
  batching:
    byte_size: 0
    condition:
      type: static
      static: false
    count: 1
    period: ""
  healthcheck: true
  id: ${!count:elastic_ids}-${!timestamp_unix}
  index: benthos_index
  max_in_flight: 1
  max_retries: 0
  pipeline: ""
  sniff: true
  timeout: 5s
  type: doc
  urls:
  - http://localhost:9200
```

Publishes messages into an Elasticsearch index. If the index does not exist then
it is created with a dynamic mapping.

Both the `id` and `index` fields can be dynamically set using function
interpolations described [here](../config_interpolation.md#functions). When
sending batched messages these interpolations are performed per message part.

### AWS Credentials

By default Benthos will use a shared credentials file when connecting to AWS
services. It's also possible to set them explicitly at the component level,
allowing you to transfer data across accounts. You can find out more
[in this document](../aws.md).

If the configured target is a managed AWS Elasticsearch cluster, you may need
to set `sniff` and `healthcheck` to false for connections to succeed.

## Fields

### `aws`

Sorry! This field is currently undocumented.

### `backoff`

Sorry! This field is currently undocumented.

### `basic_auth`

Sorry! This field is currently undocumented.

### `batching`

Sorry! This field is currently undocumented.

### `healthcheck`

Sorry! This field is currently undocumented.

### `id`

Sorry! This field is currently undocumented.

### `index`

Sorry! This field is currently undocumented.

### `max_in_flight`

Sorry! This field is currently undocumented.

### `max_retries`

Sorry! This field is currently undocumented.

### `pipeline`

Sorry! This field is currently undocumented.

### `sniff`

Sorry! This field is currently undocumented.

### `timeout`

Sorry! This field is currently undocumented.

### `type`

Sorry! This field is currently undocumented.

### `urls`

Sorry! This field is currently undocumented.

