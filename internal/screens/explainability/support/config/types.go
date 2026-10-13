package config

const (
	explainabilityConfigEnvVar      = "LONG_TOOLKIT_EXPLAINABILITY_CONFIG"
	defaultExplainabilityConfigPath = "config/explainability.json"
)

type ToolConfig struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	NodeSummary      string `json:"node_summary"`
	Subtitle         string `json:"subtitle"`
	TableTitle       string `json:"table_title"`
	SelectionTitle   string `json:"selection_title"`
	CatalogOperation string `json:"catalog_operation"`
	RunOperation     string `json:"run_operation"`
	ListCommand      string `json:"list_command"`
}

type Config struct {
	RootSummary string       `json:"root_summary"`
	Tools       []ToolConfig `json:"tools"`
}

var currentConfig Config

func cloneConfig(cfg Config) Config {
	copied := cfg
	copied.Tools = append([]ToolConfig(nil), cfg.Tools...)
	return copied
}
