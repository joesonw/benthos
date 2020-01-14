---
title: drop_on_error
type: output
---

```yaml
drop_on_error: {}
```

Attempts to write messages to a child output and if the write fails for any
reason the message is dropped instead of being reattempted. This output can be
combined with a child `retry` output in order to set an explicit
number of retry attempts before dropping a message.

For example, the following configuration attempts to send to a hypothetical
output type `foo` three times, but if all three attempts fail the
message is dropped entirely:

``` yaml
output:
  drop_on_error:
    retry:
      max_retries: 2
      output:
        type: foo
```

