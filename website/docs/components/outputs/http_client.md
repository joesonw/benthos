---
title: http_client
type: output
---

```yaml
http_client:
  backoff_on:
  - 429
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
  copy_response_headers: false
  drop_on: []
  headers:
    Content-Type: application/octet-stream
  max_in_flight: 1
  max_retry_backoff: 300s
  oauth:
    access_token: ""
    access_token_secret: ""
    consumer_key: ""
    consumer_secret: ""
    enabled: false
    request_url: ""
  propagate_response: false
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

Sends messages to an HTTP server. The request will be retried for each message
whenever the response code is outside the range of 200 -> 299 inclusive. It is
possible to list codes outside of this range in the `drop_on` field in
order to prevent retry attempts.

The period of time between retries is linear by default. Response codes that are
within the `backoff_on` list will instead apply exponential backoff
between retry attempts.

When the number of retries expires the output will reject the message, the
behaviour after this will depend on the pipeline but usually this simply means
the send is attempted again until successful whilst applying back pressure.

The URL and header values of this type can be dynamically set using function
interpolations described [here](../config_interpolation.md#functions).

The body of the HTTP request is the raw contents of the message payload. If the
message has multiple parts the request will be sent according to
[RFC1341](https://www.w3.org/Protocols/rfc1341/7_2_Multipart.html)

### Propagating Responses

It's possible to propagate the response from each HTTP request back to the input
source by setting `propagate_response` to `true`. Only inputs that
support [synchronous responses](../sync_responses.md) are able to make use of
these propagated responses.

