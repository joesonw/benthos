---
title: redis_pubsub
type: input
---

```yaml
redis_pubsub:
  channels:
  - benthos_chan
  url: tcp://localhost:6379
  use_patterns: false
```

Redis supports a publish/subscribe model, it's possible to subscribe to multiple
channels using this input.

In order to subscribe to channels using the `PSUBSCRIBE` command set
the field `use_patterns` to `true`, then you can include glob-style
patterns in your channel names. For example:

- `h?llo` subscribes to hello, hallo and hxllo
- `h*llo` subscribes to hllo and heeeello
- `h[ae]llo` subscribes to hello and hallo, but not hillo

Use `\` to escape special characters if you want to match them
verbatim.

## Fields

### `channels`

Sorry! This field is currently undocumented.

### `url`

Sorry! This field is currently undocumented.

### `use_patterns`

Sorry! This field is currently undocumented.

