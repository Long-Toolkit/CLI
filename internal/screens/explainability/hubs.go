package explainability

import (
	"errors"
	"strings"

	"github.com/charmbracelet/huh"

	toolsmodule "long-toolkit-cli/internal/screens/explainability/tools"
	"long-toolkit-cli/internal/shared/components"
)

func runExplainabilityHub() error {
	tools := Tools()
	if len(tools) == 0 {
		return errors.New("no explainability tools were found")
	}

	options := make([]huh.Option[string], 0, len(tools))
	for _, tool := range tools {
		label := strings.TrimSpace(tool.Name)
		if label == "" {
			label = "Unnamed tool"
		}
		options = append(options, huh.NewOption(label, tool.ID))
	}

	choice := tools[0].ID
	form := components.NewToolkitForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Explainability tools").
				Options(options...).
				Value(&choice),
		),
	)

	if err := form.Run(); err != nil {
		if toolsmodule.IsFlowCancelled(err) {
			return nil
		}
		return err
	}

	tool, ok := ToolByID(choice)
	if !ok {
		return errors.New("no explainability tool selected")
	}
	return runExplainabilityToolFlow(tool, false)
}
