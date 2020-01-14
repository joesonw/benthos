---
title: dynamic
type: output
---

```yaml
dynamic:
  outputs: {}
  prefix: ""
  timeout: 5s
```

The dynamic type is a special broker type where the outputs are identified by
unique labels and can be created, changed and removed during runtime via a REST
HTTP interface. The broker pattern used is always `fan_out`, meaning
each message will be delivered to each dynamic output.

To GET a JSON map of output identifiers with their current uptimes use the
'/outputs' endpoint.

To perform CRUD actions on the outputs themselves use POST, DELETE, and GET
methods on the `/outputs/{output_id}` endpoint. When using POST the
body of the request should be a JSON configuration for the output, if the output
already exists it will be changed.

## Fields

### `outputs`

Sorry! This field is currently undocumented.

### `prefix`

Sorry! This field is currently undocumented.

### `timeout`

Sorry! This field is currently undocumented.

