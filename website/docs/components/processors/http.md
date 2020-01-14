---
title: http
type: processor
---

```yaml
http:
  max_parallel: 0
  parallel: false
  request:
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
    rate_limit: ""
    retries: 3
    retry_period: 1s
    successful_on: []
    timeout: 5s
    tls:
      client_certs: []
      enabled: false
      root_cas_file: ""
      skip_cert_verify: false
    url: http://localhost:4195/post
    verb: POST
```

Performs an HTTP request using a message batch as the request body, and replaces
the original message parts with the body of the response.

If the batch contains only a single message part then it will be sent as the
body of the request. If the batch contains multiple messages then they will be
sent as a multipart HTTP request using a `Content-Type: multipart`
header.

If you are sending batches and wish to avoid this behaviour then you can set the
`parallel` flag to `true` and the messages of a batch will
be sent as individual requests in parallel. You can also cap the max number of
parallel requests with `max_parallel`. Alternatively, you can use the
[`archive`](#archive) processor to create a single message
from the batch.

The `rate_limit` field can be used to specify a rate limit
[resource](../rate_limits/README.md) to cap the rate of requests across all
parallel components service wide.

The URL and header values of this type can be dynamically set using function
interpolations described [here](../config_interpolation.md#functions).

In order to map or encode the payload to a specific request body, and map the
response back into the original payload instead of replacing it entirely, you
can use the [`process_map`](#process_map) or
 [`process_field`](#process_field) processors.

### Response Codes

Benthos considers any response code between 200 and 299 inclusive to indicate a
successful response, you can add more success status codes with the field
`successful_on`.

When a request returns a response code within the `backoff_on` field
it will be retried after increasing intervals.

When a request returns a response code within the `drop_on` field it
will not be reattempted and is immediately considered a failed request.

### Adding Metadata

If the request returns a response code this processor sets a metadata field
`http_status_code` on all resulting messages.

If the field `copy_response_headers` is set to `true` then any headers
in the response will also be set in the resulting message as metadata.
 
### Error Handling

When all retry attempts for a message are exhausted the processor cancels the
attempt. These failed messages will continue through the pipeline unchanged, but
can be dropped or placed in a dead letter queue according to your config, you
can read about these patterns [here](../error_handling.md).

## Fields

### `max_parallel`

Sorry! This field is currently undocumented.

### `parallel`

Sorry! This field is currently undocumented.

### `request`

Sorry! This field is currently undocumented.

