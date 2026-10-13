package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	toolkitconfig "long-toolkit-cli/config"
	"long-toolkit-cli/internal/shared/common"
)

func loadConfig() (Config, error) {
	raw, err := toolkitconfig.ReadBundledConfigFile(defaultExplainabilityConfigPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to load bundled explainability config: %w", err)
	}
	merged, err := decodeConfig(raw, defaultExplainabilityConfigPath)
	if err != nil {
		return Config{}, err
	}

	overridePath := strings.TrimSpace(os.Getenv(explainabilityConfigEnvVar))
	if overridePath == "" {
		return merged, nil
	}
	resolved, err := common.ResolveConfigPath(overridePath)
	if err != nil {
		return Config{}, err
	}
	overrideRaw, err := os.ReadFile(resolved)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read explainability config at %s: %w", resolved, err)
	}
	overrideConfig, err := decodeConfig(overrideRaw, resolved)
	if err != nil {
		return Config{}, err
	}

	if value := strings.TrimSpace(overrideConfig.RootSummary); value != "" {
		merged.RootSummary = value
	}
	merged.Tools = overrideConfig.Tools
	return merged, nil
}

func decodeConfig(raw []byte, source string) (Config, error) {
	var decoded Config
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return Config{}, fmt.Errorf("invalid JSON in %s: %w", source, err)
	}

	decoded.Tools = sanitiseTools(decoded.Tools)
	if len(decoded.Tools) == 0 {
		return Config{}, errors.New("explainability config produced zero tools; keep at least one tool")
	}
	return decoded, nil
}

func sanitiseTools(tools []ToolConfig) []ToolConfig {
	cleaned := make([]ToolConfig, 0, len(tools))
	seen := make(map[string]struct{}, len(tools))
	for _, tool := range tools {
		id := strings.TrimSpace(tool.ID)
		if id == "" {
			continue
		}
		if _, exists := seen[id]; exists {
			continue
		}
		catalogOperation := strings.TrimSpace(tool.CatalogOperation)
		runOperation := strings.TrimSpace(tool.RunOperation)
		if catalogOperation == "" || runOperation == "" {
			continue
		}

		name := nonEmpty(tool.Name, "Unnamed tool")
		desc := nonEmpty(tool.Description, name)
		nodeSummary := nonEmpty(tool.NodeSummary, desc)
		cleaned = append(cleaned, ToolConfig{
			ID:               id,
			Name:             name,
			Description:      desc,
			NodeSummary:      nodeSummary,
			Subtitle:         nonEmpty(tool.Subtitle, nodeSummary),
			TableTitle:       nonEmpty(tool.TableTitle, name),
			SelectionTitle:   nonEmpty(tool.SelectionTitle, name+" technique"),
			CatalogOperation: catalogOperation,
			RunOperation:     runOperation,
			ListCommand:      strings.TrimSpace(tool.ListCommand),
		})
		seen[id] = struct{}{}
	}
	return cleaned
}

func nonEmpty(primary string, fallback string) string {
	if trimmed := strings.TrimSpace(primary); trimmed != "" {
		return trimmed
	}
	return strings.TrimSpace(fallback)
}
