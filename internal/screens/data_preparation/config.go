package dataprep

import (
	dpconfig "long-toolkit-cli/internal/screens/data_preparation/support/config"
	"long-toolkit-cli/internal/shared/model"
)

type ToolConfig = dpconfig.ToolConfig
type Config = dpconfig.Config

func Configure() error {
	return dpconfig.Configure()
}

func CurrentConfig() Config {
	return dpconfig.CurrentConfig()
}

func ToolCommands() []model.CommandDef {
	return dpconfig.ToolCommands()
}

func Tools() []ToolConfig {
	return dpconfig.Tools()
}

func ToolByID(id string) (ToolConfig, bool) {
	return dpconfig.ToolByID(id)
}

func ToolFromPath(path []string) (ToolConfig, bool) {
	return dpconfig.ToolFromPath(path)
}
