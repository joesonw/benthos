---
title: switch
type: output
---

```yaml
switch:
  outputs: []
  retry_until_success: true
```

The switch output type allows you to configure multiple conditional output
targets by listing child outputs paired with conditions. Conditional logic is
currently applied per whole message batch. In order to multiplex per message of
a batch use the [`broker`](#broker) output with the pattern
`fan_out`.

In the following example, messages containing "foo" will be sent to both the
`foo` and `baz` outputs. Messages containing "bar" will be
sent to both the `bar` and `baz` outputs. Messages
containing both "foo" and "bar" will be sent to all three outputs. And finally,
messages that do not contain "foo" or "bar" will be sent to the `baz`
output only.

``` yaml
output:
  switch:
    retry_until_success: true
    outputs:
    - output:
        foo:
          foo_field_1: value1
      condition:
        text:
          operator: contains
          arg: foo
      fallthrough: true
    - output:
        bar:
          bar_field_1: value2
          bar_field_2: value3
      condition:
        text:
          operator: contains
          arg: bar
      fallthrough: true
    - output:
        baz:
          baz_field_1: value4
        processors:
        - type: baz_processor
  processors:
  - type: some_processor
```

The switch output requires a minimum of two outputs. If no condition is defined
for an output, it behaves like a static `true` condition. If
`fallthrough` is set to `true`, the switch output will
continue evaluating additional outputs after finding a match.

Messages that do not match any outputs will be dropped. If an output applies
back pressure it will block all subsequent messages.

If an output fails to send a message it will be retried continuously until
completion or service shut down. You can change this behaviour so that when an
output returns an error the switch output also returns an error by setting
`retry_until_success` to `false`. This allows you to
wrap the switch with a `try` broker, but care must be taken to ensure
duplicate messages aren't introduced during error conditions.

