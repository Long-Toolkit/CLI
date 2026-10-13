package datapp

import (
	"fmt"
	"strings"

	"long-toolkit-cli/internal/shared/components"
	"long-toolkit-cli/internal/shared/model"
	"long-toolkit-cli/internal/shared/theme"
)

func HandleBridgePath(path []string) (bool, error) {
	pathWithoutHelp, helpRequested := StripHelpFlags(path)

	canonical, handled := CanonicalPath(pathWithoutHelp)
	if !handled {
		return false, nil
	}
	if helpRequested {
		return true, renderDataPreprocessingHelp(canonical)
	}

	if tool, ok := ToolFromPath(canonical); ok {
		listOnly := strings.TrimSpace(tool.ListCommand) != "" && ContainsToken(canonical, tool.ListCommand)
		return true, runDataPreprocessingToolFlow(tool, listOnly)
	}

	if len(canonical) == 1 {
		return true, runDataPreprocessingHub()
	}
	return true, fmt.Errorf(
		"unsupported data preprocessing route in Go CLI: %s",
		strings.Join(path, " "),
	)
}

func renderDataPreprocessingHelp(canonical []string) error {
	if tool, ok := ToolFromPath(canonical); ok {
		return runDataPreprocessingToolFlow(tool, true)
	}

	cfg := CurrentConfig()

	components.PrintBlankLine()
	components.PrintLine(theme.App.SubtitleStyle().Render("Data preprocessing tools"))
	components.PrintLine(theme.App.MutedTextStyle().Render(cfg.RootSummary))
	components.PrintBlankLine()
	for _, command := range ToolCommands() {
		components.PrintLine("- " + model.CommandLabel(command) + ": " + command.Description)
	}
	return nil
}

func LoadNode(path []string) (*parsedHelp, bool, error) {
	canonical, handled := CanonicalPath(path)
	if !handled {
		return nil, false, nil
	}

	usage := strings.Join(append([]string{usageCommandPrefix}, canonical...), " ")

	newNode := func(summary string, commands []commandDef) *parsedHelp {
		return &parsedHelp{
			Path:     model.ClonePath(path),
			Usage:    usage,
			Summary:  summary,
			Commands: commands,
			Raw:      "",
		}
	}

	if len(canonical) == 1 {
		return newNode(CurrentConfig().RootSummary, ToolCommands()), true, nil
	}

	if tool, ok := ToolFromPath(canonical); ok {
		summary := strings.TrimSpace(tool.NodeSummary)
		if summary == "" {
			summary = strings.TrimSpace(tool.Description)
		}
		return newNode(summary, nil), true, nil
	}

	return nil, true, fmt.Errorf("unknown data preprocessing path: %s", strings.Join(path, " "))
}
