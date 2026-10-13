package machinelearning

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"

	toolsmodule "long-toolkit-cli/internal/screens/machine_learning/tools"
	"long-toolkit-cli/internal/shared/components"
	"long-toolkit-cli/internal/shared/theme"
)

func runMachineLearningHub() error {
	tools := Tools()
	if len(tools) == 0 {
		return errors.New("no machine learning tools were found")
	}

	options := make([]huh.Option[string], 0, len(tools))
	for _, tool := range tools {
		label := strings.TrimSpace(tool.Name)
		if label == "" {
			label = "Unnamed tool"
		}
		if !ToolIsRunnable(tool) {
			label = theme.App.MutedTextStyle().Render(
				fmt.Sprintf("%s (%s)", label, ToolStatusLabel(tool)),
			)
		}
		options = append(options, huh.NewOption(label, tool.ID))
	}

	choice := tools[0].ID
	form := components.NewToolkitForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Machine learning tools").
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
		return errors.New("no machine learning tool selected")
	}
	return runMachineLearningToolFlow(tool, false)
}
