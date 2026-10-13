package machinelearning

import (
	"strings"

	config "long-toolkit-cli/internal/screens/machine_learning/support/config"
	"long-toolkit-cli/internal/shared/model"
)

type commandDef = model.CommandDef
type parsedHelp = model.ParsedHelp

type ToolConfig = config.ToolConfig

func Configure() error {
	return config.Configure()
}

func CurrentConfig() config.Config {
	return config.CurrentConfig()
}

func ToolCommands() []model.CommandDef {
	return config.ToolCommands()
}

func Tools() []ToolConfig {
	return config.Tools()
}

func ToolByID(id string) (ToolConfig, bool) {
	return config.ToolByID(id)
}

func ToolFromPath(path []string) (ToolConfig, bool) {
	return config.ToolFromPath(path)
}

func ToolStatusLabel(tool ToolConfig) string {
	return config.StatusLabel(tool.Status)
}

func ToolIsRunnable(tool ToolConfig) bool {
	return config.IsRunnableStatus(tool.Status) &&
		strings.TrimSpace(tool.CatalogOperation) != "" &&
		strings.TrimSpace(tool.RunOperation) != ""
}
