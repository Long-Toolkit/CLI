package explainability

import (
	explconfig "long-toolkit-cli/internal/screens/explainability/support/config"
	"long-toolkit-cli/internal/shared/model"
)

type ToolConfig = explconfig.ToolConfig
type Config = explconfig.Config

func Configure() error {
	return explconfig.Configure()
}

func CurrentConfig() Config {
	return explconfig.CurrentConfig()
}

func ToolCommands() []model.CommandDef {
	return explconfig.ToolCommands()
}

func Tools() []ToolConfig {
	return explconfig.Tools()
}

func ToolByID(id string) (ToolConfig, bool) {
	return explconfig.ToolByID(id)
}

func ToolFromPath(path []string) (ToolConfig, bool) {
	return explconfig.ToolFromPath(path)
}
