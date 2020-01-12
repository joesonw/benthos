package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/Jeffail/benthos/v3/lib/buffer"
	"github.com/Jeffail/benthos/v3/lib/cache"
	"github.com/Jeffail/benthos/v3/lib/condition"
	"github.com/Jeffail/benthos/v3/lib/input"
	"github.com/Jeffail/benthos/v3/lib/metrics"
	"github.com/Jeffail/benthos/v3/lib/output"
	"github.com/Jeffail/benthos/v3/lib/processor"
	"github.com/Jeffail/benthos/v3/lib/ratelimit"
	"github.com/Jeffail/benthos/v3/lib/tracer"
	"github.com/Jeffail/benthos/v3/lib/util/config"
	"github.com/Jeffail/benthos/v3/lib/x/docs"
)

//------------------------------------------------------------------------------

func create(t, path string, resBytes []byte) {
	if err := ioutil.WriteFile(path, resBytes, 0644); err != nil {
		panic(err)
	}
	fmt.Printf("Generated '%v' doc at: %v\n", t, path)
}

func render(dir string, confSanit interface{}, spec docs.ComponentSpec) {
	var segment interface{}
	switch t := confSanit.(type) {
	case map[string]interface{}:
		segment = t[spec.Name]
	case config.Sanitised:
		segment = t[spec.Name]
	default:
		panic(fmt.Sprintf("Failed to generate docs for '%v': sanitised config wrong type: %T", spec.Name, confSanit))
	}

	mdSpec, err := spec.AsMarkdown(segment)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate docs for '%v': %v", spec.Name, err))
	}

	create(spec.Name, dir, mdSpec)
}

func main() {
	docsDir := "./website/docs/components"
	flag.StringVar(&docsDir, "dir", docsDir, "The directory to write docs to")
	flag.Parse()

	doInputs(docsDir)
	doBuffers(docsDir)
	doCaches(docsDir)
	doConditions(docsDir)
	doMetrics(docsDir)
	doOutputs(docsDir)
	doProcessors(docsDir)
	doRateLimits(docsDir)
	doTracers(docsDir)
}

func doInputs(docsDir string) {
	for k, v := range input.Constructors {
		if v.Deprecated {
			continue
		}
		spec := docs.ComponentSpec{
			Type:        "input",
			Name:        k,
			Description: v.Description,
			Fields:      v.FieldSpecs,
		}

		conf := input.NewConfig()
		conf.Type = k
		confSanit, err := input.SanitiseConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate docs for '%v': %v", k, err))
		}

		render(path.Join(docsDir, "./inputs", k+".md"), confSanit, spec)
	}
}

func doBuffers(docsDir string) {
	for k, v := range buffer.Constructors {
		spec := docs.ComponentSpec{
			Type:        "buffer",
			Name:        k,
			Description: v.Description,
			Fields:      v.FieldSpecs,
		}

		conf := buffer.NewConfig()
		conf.Type = k
		confSanit, err := buffer.SanitiseConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate docs for '%v': %v", k, err))
		}

		render(path.Join(docsDir, "./buffers", k+".md"), confSanit, spec)
	}
}

func doCaches(docsDir string) {
	for k, v := range cache.Constructors {
		spec := docs.ComponentSpec{
			Type:        "cache",
			Name:        k,
			Description: v.Description,
			Fields:      v.FieldSpecs,
		}

		conf := cache.NewConfig()
		conf.Type = k
		confSanit, err := cache.SanitiseConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate docs for '%v': %v", k, err))
		}

		render(path.Join(docsDir, "./caches", k+".md"), confSanit, spec)
	}
}

func doConditions(docsDir string) {
	for k, v := range condition.Constructors {
		spec := docs.ComponentSpec{
			Type:        "condition",
			Name:        k,
			Description: v.Description,
			Fields:      v.FieldSpecs,
		}

		conf := condition.NewConfig()
		conf.Type = k
		confSanit, err := condition.SanitiseConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate docs for '%v': %v", k, err))
		}

		render(path.Join(docsDir, "./conditions", k+".md"), confSanit, spec)
	}
}

func doMetrics(docsDir string) {
	for k, v := range metrics.Constructors {
		spec := docs.ComponentSpec{
			Type:        "metrics",
			Name:        k,
			Description: v.Description,
			Fields:      v.FieldSpecs,
		}

		conf := metrics.NewConfig()
		conf.Type = k
		confSanit, err := metrics.SanitiseConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate docs for '%v': %v", k, err))
		}

		render(path.Join(docsDir, "./metrics", k+".md"), confSanit, spec)
	}
}

func doOutputs(docsDir string) {
	for k, v := range output.Constructors {
		spec := docs.ComponentSpec{
			Type:        "output",
			Name:        k,
			Description: v.Description,
			Fields:      v.FieldSpecs,
		}

		conf := output.NewConfig()
		conf.Type = k
		confSanit, err := output.SanitiseConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate docs for '%v': %v", k, err))
		}

		render(path.Join(docsDir, "./outputs", k+".md"), confSanit, spec)
	}
}

func doProcessors(docsDir string) {
	for k, v := range processor.Constructors {
		spec := docs.ComponentSpec{
			Type:        "processor",
			Name:        k,
			Description: v.Description,
			Fields:      v.FieldSpecs,
		}

		conf := processor.NewConfig()
		conf.Type = k
		confSanit, err := processor.SanitiseConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate docs for '%v': %v", k, err))
		}

		render(path.Join(docsDir, "./processors", k+".md"), confSanit, spec)
	}
}

func doRateLimits(docsDir string) {
	for k, v := range ratelimit.Constructors {
		spec := docs.ComponentSpec{
			Type:        "rate_limit",
			Name:        k,
			Description: v.Description,
			Fields:      v.FieldSpecs,
		}

		conf := ratelimit.NewConfig()
		conf.Type = k
		confSanit, err := ratelimit.SanitiseConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate docs for '%v': %v", k, err))
		}

		render(path.Join(docsDir, "./rate_limits", k+".md"), confSanit, spec)
	}
}

func doTracers(docsDir string) {
	for k, v := range tracer.Constructors {
		spec := docs.ComponentSpec{
			Type:        "tracer",
			Name:        k,
			Description: v.Description,
			Fields:      v.FieldSpecs,
		}

		conf := tracer.NewConfig()
		conf.Type = k
		confSanit, err := tracer.SanitiseConfig(conf)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate docs for '%v': %v", k, err))
		}

		render(path.Join(docsDir, "./tracers", k+".md"), confSanit, spec)
	}
}

//------------------------------------------------------------------------------
