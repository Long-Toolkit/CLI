package tools

import (
	"strings"

	aggregatelongtocrosssectional "long-toolkit-cli/internal/screens/data_preprocessing/tools/aggregate_long_to_cross_sectional"
	cleandataset "long-toolkit-cli/internal/screens/data_preprocessing/tools/clean_dataset"
	harmonisecategories "long-toolkit-cli/internal/screens/data_preprocessing/tools/harmonise_categories"
	missingimputation "long-toolkit-cli/internal/screens/data_preprocessing/tools/missing_imputation"
	pivotlongtowide "long-toolkit-cli/internal/screens/data_preprocessing/tools/pivot_long_to_wide"
	"long-toolkit-cli/internal/screens/data_preprocessing/tools/remove_columns"
	renamefeature "long-toolkit-cli/internal/screens/data_preprocessing/tools/rename_feature"
	showtable "long-toolkit-cli/internal/screens/data_preprocessing/tools/show_table"
	"long-toolkit-cli/internal/screens/data_preprocessing/tools/spec"
)

type Tool = spec.Tool
type FlowSpec = spec.FlowSpec

type Builder func(tool Tool) FlowSpec

var builders = map[string]Builder{
	"aggregate_long_to_cross_sectional": aggregatelongtocrosssectional.Build,
	"clean_dataset":                     cleandataset.Build,
	"harmonise_categories":              harmonisecategories.Build,
	"missing_imputation":                missingimputation.Build,
	"pivot_long_to_wide":                pivotlongtowide.Build,
	"remove_columns":                    removecolumns.Build,
	"rename_feature":                    renamefeature.Build,
	"show_table":                        showtable.Build,
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
