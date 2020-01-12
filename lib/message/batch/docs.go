package batch

import "github.com/Jeffail/benthos/v3/lib/x/docs"

// FieldSpec returns a spec for a common batching field.
func FieldSpec() docs.FieldSpec {
	return docs.FieldSpec{
		Description: `
Allows you to configure a [batching policy](/docs/configuration/batching).`,
		Examples: []interface{}{
			map[string]interface{}{
				"byte_size": 5000,
				"period":    "1s",
			},
			map[string]interface{}{
				"count":  10,
				"period": "1s",
			},
			map[string]interface{}{
				"period": "1m",
				"condition": map[string]interface{}{
					"text": map[string]interface{}{
						"operator": "contains",
						"arg":      "END BATCH",
					},
				},
			},
		},
	}
}
