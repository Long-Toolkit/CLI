package tools

import (
	"strings"

	shapanalysis "long-toolkit-cli/internal/screens/explainability/tools/shap_analysis"
	"long-toolkit-cli/internal/screens/machine_learning/tools/spec"
)

type Tool = spec.Tool
type FlowSpec = spec.FlowSpec

type Builder func(tool Tool) FlowSpec

var builders = map[string]Builder{
	"shap_analysis": shapanalysis.Build,
}

func Resolve(tool Tool) (FlowSpec, bool) {
	builder, ok := builders[canonicalToolKey(tool.ID)]
	if !ok {
		return FlowSpec{}, false
	}
	return builder(tool), true
}

func Default(tool Tool) FlowSpec {
	return spec.Default(tool)
}

func canonicalToolKey(raw string) string {
	return strings.TrimSpace(strings.ToLower(raw))
}
