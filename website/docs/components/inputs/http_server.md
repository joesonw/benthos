---
title: http_server
type: input
---

```yaml
http_server:
  address: ""
  cert_file: ""
  key_file: ""
  path: /post
  rate_limit: ""
  sync_response:
    headers:
      Content-Type: application/octet-stream
  timeout: 5s
  ws_path: /post/ws
  ws_rate_limit_message: ""
  ws_welcome_message: ""
```

Receive messages POSTed over HTTP(S). HTTP 2.0 is supported when using TLS,
which is enabled when key and cert files are specified.

You can leave the 'address' config field blank in order to use the instance wide
HTTP server.

The field `rate_limit` allows you to specify an optional
[`rate_limit` resource](/docs/components/rate_limits/about), which will be
applied to each HTTP request made and each websocket payload received.

When the rate limit is breached HTTP requests will have a 429 response returned
with a Retry-After header. Websocket payloads will be dropped and an optional
response payload will be sent as per `ws_rate_limit_message`.

### Responses

It's possible to return a response for each message received using
[synchronous responses](/docs/guides/sync_responses). When doing so you can customise
headers with the `sync_response` field `headers`, which can also use
[function interpolation](/docs/configuration/interpolation#metadata) in the value based
on the response message contents.

### Endpoints

The following fields specify endpoints that are registered for sending messages:

#### `path` (defaults to `/post`)

This endpoint expects POST requests where the entire request body is consumed as
a single message.

If the request contains a multipart `content-type` header as per
[rfc1341](https://www.w3.org/Protocols/rfc1341/7_2_Multipart.html) then the
multiple parts are consumed as a batch of messages, where each body part is a
message of the batch.

#### `ws_path` (defaults to `/post/ws`)

Creates a websocket connection, where payloads received on the socket are passed
through the pipeline as a batch of one message.

You may specify an optional `ws_welcome_message`, which is a static
payload to be sent to all clients once a websocket connection is first
established.

It's also possible to specify a `ws_rate_limit_message`, which is a
static payload to be sent to clients that have triggered the servers rate limit.

### Metadata

This input adds the following metadata fields to each message:

``` text
- http_server_user_agent
- All headers (only first values are taken)
- All query parameters
- All cookies
```

You can access these metadata fields using
[function interpolation](/docs/configuration/interpolation#metadata).

