---
title: try
type: output
---

```yaml
try: []
```

Attempts to send each message to only one output, starting from the first output
on the list. If an output attempt fails then the next output in the list is
attempted, and so on.

This pattern is useful for triggering events in the case where certain output
targets have broken. For example, if you had an output type `http_client`
but wished to reroute messages whenever the endpoint becomes unreachable you
could use this pattern:

``` yaml
output:
  try:
  - http_client:
      url: http://foo:4195/post/might/become/unreachable
      retries: 3
      retry_period: 1s
  - http_client:
      url: http://bar:4196/somewhere/else
      retries: 3
      retry_period: 1s
    processors:
    - text:
        operator: prepend
        value: 'failed to send this message to foo: '
  - file:
      path: /usr/local/benthos/everything_failed.jsonl
```

