---
title: json
type: processor
---

```yaml
json:
  operator: clean
  parts: []
  path: ""
  value: ""
```

Parses messages as a JSON document, performs a mutation on the data, and then
overwrites the previous contents with the new value.

The field `path` is a [dot separated path](../field_paths.md) which,
for most operators, determines the field within the payload to be targeted. If
the path is empty or "." the root of the data will be targeted.

This processor will interpolate functions within the 'value' field, you can find
a list of functions [here](../config_interpolation.md#functions).

### Operators

`append`

Appends a value to an array at a target dot path. If the path does not exist all
objects in the path are created (unless there is a collision).

If a non-array value already exists in the target path it will be replaced by an
array containing the original value as well as the new value.

If the value is an array the elements of the array are expanded into the new
array. E.g. if the target is an array `[0,1]` and the value is also an
array `[2,3]`, the result will be `[0,1,2,3]` as opposed to
`[0,1,[2,3]]`.

`clean`

Walks the JSON structure and deletes any fields where the value is:

- An empty array
- An empty object
- An empty string
- null

`copy`

Copies the value of a target dot path (if it exists) to a location. The
destination path is specified in the `value` field. If the destination
path does not exist all objects in the path are created (unless there is a
collision).

`delete`

Removes a key identified by the dot path. If the path does not exist this is a
no-op.

`explode`

Explodes an array or object within a JSON document.

Exploding arrays results in a root level array containing elements matching the
original document, where the target field of each element is an element of the
exploded array.

Exploding objects results in a root level object where the keys match the target
object, and the values match the original document but with the target field
replaced by the exploded value.

It is then possible to expand the result to create individual messages per
element with the [`unarchive` processor](#unarchive) `json_array` or
`json_object` format..

For example, given the following config:

```yaml
json:
  operator: explode
  path: value
```

And two input documents:

```json
{"id":1,"value":["foo","bar","baz"]}
{"id":1,"value":{"foo":2,"bar":[3,4],"baz":{"bev":5}}}
```

The respective results would be:

```json
[{"id":1,"value":"foo"},{"id":1,"value":"bar"},{"id":1,"value":"baz"}]
{"foo":{"id":1,"value":2},"bar":{"id":1,"value":[3,4]},"baz":{"id":1,"value":{"bev":5}}}
```

`move`

Moves the value of a target dot path (if it exists) to a new location. The
destination path is specified in the `value` field. If the destination
path does not exist all objects in the path are created (unless there is a
collision).

`select`

Reads the value found at a dot path and replaced the original contents entirely
by the new value.

`set`

Sets the value of a field at a dot path. If the path does not exist all objects
in the path are created (unless there is a collision).

The value can be any type, including objects and arrays. When using YAML
configuration files a YAML object will be converted into a JSON object, i.e.
with the config:

```yaml
json:
  operator: set
  parts: [0]
  path: some.path
  value:
    foo:
      bar: 5
```

The value will be converted into '{"foo":{"bar":5}}'. If the YAML object
contains keys that aren't strings those fields will be ignored.

`split`

Splits a string field by a value and replaces the original string with an array
containing the results of the split. This operator requires both the path value
and the contents of the `value` field to be strings.

