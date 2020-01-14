---
title: sync_response
type: processor
---

```yaml
sync_response: {}
```

Adds the payload in its current state as a synchronous response to the input
source, where it is dealt with according to that specific input type.

For most inputs this mechanism is ignored entirely, in which case the sync
response is dropped without penalty. It is therefore safe to use this processor
even when combining input types that might not have support for sync responses.
An example of an input able to utilise this is the `http_server`.

For more information please read [Synchronous Responses](../sync_responses.md).

