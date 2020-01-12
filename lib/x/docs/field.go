package docs

import "gopkg.in/yaml.v3"

//------------------------------------------------------------------------------

// FieldInterpolation represents a type of interpolation supported by a field.
type FieldInterpolation int

// Interpolation types.
const (
	FieldInterpolationNone FieldInterpolation = iota
	FieldInterpolationBatchWide
	FieldInterpolationIndividual
)

// FieldSpec describes a component config field.
type FieldSpec struct {
	// Name of the field (as it appears in config).
	Name string

	// Description of the field purpose (in markdown).
	Description string

	// Advanced is true for optional fields that will not be present in most
	// configs.
	Advanced bool

	// Deprecated is true for fields that are deprecated and only exist for
	// backwards compatibility reasons.
	Deprecated bool

	// Type of the field. This is optional and doesn't prevent documentation for
	// a field.
	Type string

	// Interpolation indicates the type of interpolation that this field
	// supports.
	Interpolation FieldInterpolation

	// Examples is a slice of optional example values for a field.
	Examples []interface{}

	// Children fields of this field (it must be an object).
	Children FieldSpecs
}

// SupportsInterpolation returns a new FieldSpec that specifies whether it
// supports function interpolation (batch wide or not).
func (f FieldSpec) SupportsInterpolation(batchWide bool) FieldSpec {
	if batchWide {
		f.Interpolation = FieldInterpolationBatchWide
	} else {
		f.Interpolation = FieldInterpolationIndividual
	}
	return f
}

// HasType returns a new FieldSpec that specifies a specific type.
func (f FieldSpec) HasType(typeStr string) FieldSpec {
	f.Type = typeStr
	return f
}

// WithChildren returns a new FieldSpec that has child fields.
func (f FieldSpec) WithChildren(children FieldSpecs) FieldSpec {
	f.Type = "object"
	f.Children = children
	return f
}

// FieldAdvanced returns a field spec for an advanced field.
func FieldAdvanced(name, description string, examples ...interface{}) FieldSpec {
	return FieldSpec{
		Name:        name,
		Description: description,
		Advanced:    true,
		Examples:    examples,
	}
}

// FieldCommon returns a field spec for a common field.
func FieldCommon(name, description string, examples ...interface{}) FieldSpec {
	return FieldSpec{
		Name:        name,
		Description: description,
		Examples:    examples,
	}
}

// FieldDeprecated returns a field spec for a deprecated field.
func FieldDeprecated(name string) FieldSpec {
	return FieldSpec{
		Name:        name,
		Description: "DEPRECATED: Do not use.",
		Deprecated:  true,
	}
}

//------------------------------------------------------------------------------

// FieldSpecs is a slice of field specs for a component.
type FieldSpecs []FieldSpec

// ConfigCommon takes a sanitised configuration of a component, a map of field
// docs, and removes all fields that aren't common or are deprecated.
func (f FieldSpecs) ConfigCommon(config interface{}) (map[string]interface{}, error) {
	return f.configFiltered(config, func(f FieldSpec) bool {
		return !(f.Advanced || f.Deprecated)
	})
}

// ConfigAdvanced takes a sanitised configuration of a component, a map of field
// docs, and removes all fields that are deprecated.
func (f FieldSpecs) ConfigAdvanced(config interface{}) (map[string]interface{}, error) {
	return f.configFiltered(config, func(f FieldSpec) bool {
		return !f.Deprecated
	})
}

func (f FieldSpecs) configFiltered(config interface{}, filter func(f FieldSpec) bool) (map[string]interface{}, error) {
	var asMap map[string]interface{}
	var ok bool
	if asMap, ok = config.(map[string]interface{}); !ok {
		rawBytes, err := yaml.Marshal(config)
		if err != nil {
			return nil, err
		}
		if err = yaml.Unmarshal(rawBytes, &asMap); err != nil {
			return nil, err
		}
	}

	newMap := map[string]interface{}{}
	acceptedKeys := map[string]FieldSpec{}
	for _, v := range f {
		if filter(v) {
			acceptedKeys[v.Name] = v
		}
	}
	for k, v := range asMap {
		if field, exists := acceptedKeys[k]; exists {
			if len(field.Children) > 0 {
				var err error
				if v, err = field.Children.configFiltered(v, filter); err != nil {
					return nil, err
				}
			}
			newMap[k] = v
		}
	}
	return newMap, nil
}

//------------------------------------------------------------------------------
