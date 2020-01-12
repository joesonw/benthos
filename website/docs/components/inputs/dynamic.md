---
title: dynamic
type: input
---

```yaml
dynamic:
  inputs: {}
  prefix: ""
  timeout: 5s
```

The dynamic type is a special broker type where the inputs are identified by
unique labels and can be created, changed and removed during runtime via a REST
HTTP interface.

To GET a JSON map of input identifiers with their current uptimes use the
`/inputs` endpoint.

To perform CRUD actions on the inputs themselves use POST, DELETE, and GET
methods on the `/inputs/{input_id}` endpoint. When using POST the body
of the request should be a JSON configuration for the input, if the input
already exists it will be changed.

## Fields

### `inputs`

Sorry! This field is currently undocumented.

### `prefix`

Sorry! This field is currently undocumented.

### `timeout`

Sorry! This field is currently undocumented.

