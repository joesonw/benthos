package docs

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"

	"github.com/Jeffail/benthos/v3/lib/util/config"
	"github.com/Jeffail/gabs/v2"
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

` + "`{{$field.Type}}`" + ` {{$field.Description}}
{{if gt (len $field.Examples) 0 -}}
` + "```yaml" + `
# Examples

{{range $j, $example := $field.Examples -}}
{{$example}}
{{end -}}
` + "```" + `
{{end -}}
{{end -}}
`

func (c *ComponentSpec) createConfigs(fullConfigExample interface{}) (
	advancedConfigBytes, commonConfigBytes []byte,
	advancedConfig interface{},
) {
	var err error
	if len(c.Fields) > 0 {
		advancedConfig, err = c.Fields.ConfigAdvanced(fullConfigExample)
		if err == nil {
			advancedConfigBytes, err = config.MarshalYAML(map[string]interface{}{
				c.Name: advancedConfig,
			})
		}
		if err == nil {
			advancedConfig = fullConfigExample
		}
		var commonConfig interface{}
		if err == nil {
			commonConfig, err = c.Fields.ConfigCommon(fullConfigExample)
		}
		if err == nil {
			commonConfigBytes, err = config.MarshalYAML(map[string]interface{}{
				c.Name: commonConfig,
			})
		}
	}
	if err != nil {
		panic(err)
	}
	if err != nil || len(c.Fields) == 0 {
		if advancedConfigBytes, err = config.MarshalYAML(map[string]interface{}{
			c.Name: fullConfigExample,
		}); err != nil {
			panic(err)
		}
		commonConfigBytes = advancedConfigBytes
		advancedConfig = fullConfigExample
	}
	return
}

// AsMarkdown renders the spec of a component, along with a full configuration
// example, into a markdown document.
func (c *ComponentSpec) AsMarkdown(fullConfigExample interface{}) ([]byte, error) {
	ctx := componentContext{
		Name:        c.Name,
		Type:        c.Type,
		Description: c.Description,
	}

	advancedConfigBytes, commonConfigBytes, advancedConfig := c.createConfigs(fullConfigExample)
	ctx.CommonConfig = string(commonConfigBytes)
	ctx.AdvancedConfig = string(advancedConfigBytes)

	gConf := gabs.Wrap(advancedConfig)

	if c.Description[0] == '\n' {
		ctx.Description = c.Description[1:]
	}

	flattenedFields := FieldSpecs{}
	var walkFields func(root string, f FieldSpecs)
	walkFields = func(root string, f FieldSpecs) {
		for _, v := range f {
			newV := v
			newV.Children = nil
			if len(root) > 0 {
				newV.Name = root + newV.Name
			}
			flattenedFields = append(flattenedFields, newV)
			if len(v.Children) > 0 {
				walkFields(v.Name+".", v.Children)
			}
		}
	}
	walkFields("", c.Fields)

	for _, v := range flattenedFields {
		if v.Deprecated {
			continue
		}

		if !gConf.ExistsP(v.Name) {
			return nil, fmt.Errorf("unrecognised field '%v'", v.Name)
		}

		fieldType := v.Type
		if len(fieldType) == 0 {
			if len(v.Examples) > 0 {
				fieldType = reflect.TypeOf(v.Examples[0]).Kind().String()
			} else {
				if c := gConf.Path(v.Name).Data(); c != nil {
					fieldType = reflect.TypeOf(c).Kind().String()
				} else {
					return nil, fmt.Errorf("unable to infer type of '%v'", v.Name)
				}
			}
		}
		switch fieldType {
		case "map":
			fieldType = "object"
		case "slice":
			fieldType = "array"
		case "float64", "int", "int64":
			fieldType = "number"
		}

		var examples []string
		for _, example := range v.Examples {
			exampleBytes, err := config.MarshalYAML(map[string]interface{}{
				v.Name: example,
			})
			if err != nil {
				return nil, err
			}
			examples = append(examples, string(exampleBytes))
		}

		fieldCtx := fieldContext{
			Name:        v.Name,
			Type:        fieldType,
			Description: v.Description,
			Advanced:    v.Advanced,
			Examples:    examples,
		}

		if len(fieldCtx.Description) == 0 {
			fieldCtx.Description = "Sorry! This field is missing documentation."
		}

		if fieldCtx.Description[0] == '\n' {
			fieldCtx.Description = fieldCtx.Description[1:]
		}

		ctx.Fields = append(ctx.Fields, fieldCtx)
	}

	var buf bytes.Buffer
	err := template.Must(template.New("component").Parse(componentTemplate)).Execute(&buf, ctx)

	return buf.Bytes(), err
}
