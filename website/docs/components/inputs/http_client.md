---
title: http_client
type: input
---

```yaml
http_client:
  backoff_on:
  - 429
  basic_auth:
    enabled: false
    password: ""
    username: ""
  copy_response_headers: false
  drop_on: []
  headers:
    Content-Type: application/octet-stream
  max_retry_backoff: 300s
  oauth:
    access_token: ""
    access_token_secret: ""
    consumer_key: ""
    consumer_secret: ""
    enabled: false
    request_url: ""
  payload: ""
  rate_limit: ""
  retries: 3
  retry_period: 1s
  stream:
    delimiter: ""
    enabled: false
    max_buffer: 1e+06
    multipart: false
    reconnect: true
  successful_on: []
  timeout: 5s
  tls:
    client_certs: []
    enabled: false
    root_cas_file: ""
    skip_cert_verify: false
  url: http://localhost:4195/get
  verb: GET
```

The HTTP client input type connects to a server and continuously performs
requests for a single message.

You should set a sensible retry period and max backoff so as to not flood your
target server.

The URL and header values of this type can be dynamically set using function
interpolations described [here](../config_interpolation.md#functions).

### Streaming

If you enable streaming then Benthos will consume the body of the response as a
line delimited list of message parts. Each part is read as an individual message
unless multipart is set to true, in which case an empty line indicates the end
of a message.

## Fields

### `backoff_on`

Sorry! This field is currently undocumented.

### `basic_auth`

Sorry! This field is currently undocumented.

### `copy_response_headers`

Sorry! This field is currently undocumented.

### `drop_on`

Sorry! This field is currently undocumented.

### `headers`

Sorry! This field is currently undocumented.

### `max_retry_backoff`

Sorry! This field is currently undocumented.

### `oauth`

Sorry! This field is currently undocumented.

### `payload`

Sorry! This field is currently undocumented.

### `rate_limit`

Sorry! This field is currently undocumented.

### `retries`

Sorry! This field is currently undocumented.

### `retry_period`

Sorry! This field is currently undocumented.

### `stream`

Sorry! This field is currently undocumented.

### `successful_on`

Sorry! This field is currently undocumented.

### `timeout`

Sorry! This field is currently undocumented.

### `tls`

Sorry! This field is currently undocumented.

### `url`

Sorry! This field is currently undocumented.

### `verb`

Sorry! This field is currently undocumented.

