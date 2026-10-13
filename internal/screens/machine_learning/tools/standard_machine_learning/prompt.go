package standardmachinelearning

import (
	"long-toolkit-cli/internal/screens/machine_learning/support/schema"
	"long-toolkit-cli/internal/screens/machine_learning/tools/prompting"
)

func PromptTechniqueSelection(title string, techniques []schema.Technique) (schema.Technique, error) {
	return prompting.PromptTechniqueSelection(title, techniques)
}

func PromptTechniqueParameters(technique schema.Technique) (map[string]any, error) {
	return prompting.PromptTechniqueParameters(technique)
}
