---
title: http_server
type: output
---

```yaml
http_server:
  address: ""
  cert_file: ""
  key_file: ""
  path: /get
  stream_path: /get/stream
  timeout: 5s
  ws_path: /get/ws
```

Sets up an HTTP server that will send messages over HTTP(S) GET requests. HTTP
2.0 is supported when using TLS, which is enabled when key and cert files are
specified.

You can leave the `address` config field blank in order to use the
default service wide server address, but this will ignore TLS options.

Three endpoints will be registered at the paths specified by the fields
`path`, `stream_path` and `ws_path`. Which allow you to consume a
single message batch, a continuous stream of line delimited messages, or a
websocket of messages for each request respectively.

When messages are batched the `path` endpoint encodes the batch
according to [RFC1341](https://www.w3.org/Protocols/rfc1341/7_2_Multipart.html).

