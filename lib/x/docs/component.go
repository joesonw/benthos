package docs

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"text/template"

	"github.com/Jeffail/benthos/v3/lib/util/config"
)

// ComponentSpec describes a Benthos component.
type ComponentSpec struct {
	// Name of the component
	Name string

	// Type of the component (input, output, etc)
	Type string

	// Description of the component (in markdown).
	Description string

	Fields FieldSpecs
}

type fieldContext struct {
	Name        string
	Type        string
	Description string
	Advanced    bool
	Deprecated  bool
	Examples    []string
}

type componentContext struct {
	Name           string
	Type           string
	Description    string
	Fields         []fieldContext
	CommonConfig   string
	AdvancedConfig string
}

/*
{{if $field.Advanced}}(Advanced){{end}}

Type: ` + "`{{$field.Type}}`" + `
*/

var componentTemplate = `---
title: {{.Name}}
type: {{.Type}}
---

{{if eq .CommonConfig .AdvancedConfig -}}
` + "```yaml" + `
{{.CommonConfig -}}
` + "```" + `
{{else}}
import Tabs from '@theme/Tabs';

<Tabs defaultValue="common" values={{"{"}}[
  { label: 'Common', value: 'common', },
  { label: 'Advanced', value: 'advanced', },
]{{"}"}}>

import TabItem from '@theme/TabItem';

<TabItem value="common">

` + "```yaml" + `
{{.CommonConfig -}}
` + "```" + `

</TabItem>
<TabItem value="advanced">

` + "```yaml" + `
{{.AdvancedConfig -}}
` + "```" + `

</TabItem>
</Tabs>
{{end}}
{{.Description}}

{{if gt (len .Fields) 0 -}}
## Fields

{{end -}}
{{range $i, $field := .Fields -}}
### ` + "`{{$field.Name}}`" + `

{{if gt (len $field.Description) 0 -}}
{{$field.Description}}
{{else -}}
Sorry! This field is currently undocumented.
{{end}}
{{if gt (len $field.Examples) 0 -}}
` + "```yaml" + `
{{range $j, $example := $field.Examples -}}
{{$example}}
{{end -}}
` + "```" + `
{{end -}}
{{end -}}
`

// AsMarkdown renders the spec of a component, along with a full configuration
// example, into a markdown document.
func (c *ComponentSpec) AsMarkdown(fullConfigExample interface{}) ([]byte, error) {
	ctx := componentContext{
		Name:        c.Name,
		Type:        c.Type,
		Description: c.Description,
	}

	var advancedConfigBytes []byte
	var commonConfigBytes []byte
	var advancedConfig map[string]interface{}
	var err error

	if asMap, isMap := fullConfigExample.(map[string]interface{}); isMap {
		advancedConfig = c.Fields.ConfigAdvanced(asMap)
		advancedConfigBytes, err = config.MarshalYAML(map[string]interface{}{
			c.Name: advancedConfig,
		})
		commonConfig := c.Fields.ConfigCommon(asMap)
		if err == nil {
			commonConfigBytes, err = config.MarshalYAML(map[string]interface{}{
				c.Name: commonConfig,
			})
		}
	} else {
		advancedConfigBytes, err = config.MarshalYAML(map[string]interface{}{
			c.Name: fullConfigExample,
		})
		commonConfigBytes = advancedConfigBytes
	}
	if err != nil {
		return nil, err
	}

	ctx.CommonConfig = string(commonConfigBytes)
	ctx.AdvancedConfig = string(advancedConfigBytes)

	if c.Description[0] == '\n' {
		ctx.Description = c.Description[1:]
	}

	fieldNames := []string{}
	unrecognisedSpecs := []string{}
	for k, spec := range c.Fields {
		if spec.Deprecated {
			continue
		}
		fieldNames = append(fieldNames, k)
		if _, exists := advancedConfig[k]; !exists {
			unrecognisedSpecs = append(unrecognisedSpecs, k)
		}
	}
	if len(unrecognisedSpecs) > 0 {
		return nil, fmt.Errorf("unrecognised fields found within '%v' spec: %v", c.Name, unrecognisedSpecs)
	}
	for k := range advancedConfig {
		if _, exists := c.Fields[k]; !exists {
			fieldNames = append(fieldNames, k)
		}
	}
	sort.Strings(fieldNames)

	for _, k := range fieldNames {
		v := c.Fields[k]
		if v.Deprecated {
			continue
		}

		fieldType := reflect.TypeOf(advancedConfig[k]).Kind().String()
		switch fieldType {
		case "map":
			fieldType = "object"
		case "slice":
			fieldType = "array"
		}

		var examples []string
		for _, example := range v.Examples {
			exampleBytes, err := config.MarshalYAML(map[string]interface{}{
				k: example,
			})
			if err != nil {
				return nil, err
			}
			examples = append(examples, string(exampleBytes))
		}

		fieldCtx := fieldContext{
			Name:        k,
			Type:        fieldType,
			Description: v.Description,
			Advanced:    v.Advanced,
			Examples:    examples,
		}

		if len(fieldCtx.Description) > 0 && fieldCtx.Description[0] == '\n' {
			fieldCtx.Description = fieldCtx.Description[1:]
		}

		ctx.Fields = append(ctx.Fields, fieldCtx)
	}

	var buf bytes.Buffer
	err = template.Must(template.New("component").Parse(componentTemplate)).Execute(&buf, ctx)

	return buf.Bytes(), err
}
