package tools

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"long-toolkit-cli/internal/screens/machine_learning/support/schema"
	"long-toolkit-cli/internal/screens/machine_learning/support/ui"
	"long-toolkit-cli/internal/screens/machine_learning/tools/prompting"
	"long-toolkit-cli/internal/screens/machine_learning/tools/spec"
	"long-toolkit-cli/internal/shared/components"
	"long-toolkit-cli/internal/shared/theme"
)

// Run drives one explainability tool: fetch its catalogue, pick a technique,
// collect parameters, call the bridge and print a summary. It reuses the
// machine learning stage's generic schema, screen and prompting helpers.
func Run(def Definition, listOnly bool, runtime Runtime) error {
	if runtime.Execute == nil {
		return errors.New("explainability tools runtime has no bridge executor")
	}

	screenTool := Tool{
		ID:             def.ID,
		Name:           def.Name,
		Subtitle:       def.Subtitle,
		TableTitle:     def.TableTitle,
		SelectionTitle: def.SelectionTitle,
	}
	flow, ok := Resolve(screenTool)
	if !ok {
		flow = Default(screenTool)
	}

	ui.PrepareActionScreen(flow.Title, flow.Subtitle, runtime.InNavigator)

	catalog, err := runtime.Execute(def.CatalogOperation, map[string]any{})
	if err != nil {
		return err
	}
	techniques, err := schema.DecodeTechniquePayload(catalog)
	if err != nil {
		return err
	}
	if len(techniques) == 0 {
		return fmt.Errorf("no techniques were found for %s", strings.ToLower(flow.Title))
	}

	if listOnly {
		ui.PrintTechniqueTable(techniques)
		return nil
	}

	selected := techniques[0]
	if len(techniques) > 1 {
		ui.PrintTechniqueTable(techniques)
		selectionTitle := strings.TrimSpace(flow.SelectionTitle)
		if selectionTitle == "" {
			selectionTitle = fmt.Sprintf("%s technique", strings.TrimSpace(flow.Title))
		}
		if flow.PromptTechniqueSelection == nil {
			flow.PromptTechniqueSelection = prompting.PromptTechniqueSelection
		}
		selected, err = flow.PromptTechniqueSelection(selectionTitle, techniques)
		if err != nil {
			return err
		}
	}

	if flow.PromptTechniqueParams == nil {
		flow.PromptTechniqueParams = prompting.PromptTechniqueParameters
	}
	params, err := flow.PromptTechniqueParams(selected)
	if err != nil {
		return err
	}

	result, err := runtime.Execute(
		def.RunOperation,
		map[string]any{
			"technique": selected.ID,
			"params":    params,
		},
	)
	if err != nil {
		return err
	}

	if flow.SummaryPrinter != nil {
		flow.SummaryPrinter(selected, result)
	} else {
		printGenericSummary(selected, result)
	}

	return ui.RunExitCountdown("Returning to toolkit", 10*time.Second, runtime.InNavigator)
}

func printGenericSummary(selected schema.Technique, result map[string]any) {
	components.PrintBlankLine()
	components.PrintLine(theme.App.SubtitleStyle().Render("Run complete"))
	components.PrintfLine("- Technique: %s", spec.NonEmpty(selected.Name, selected.ID))
	for key, value := range result {
		components.PrintfLine("- %s: %v", strings.TrimSpace(key), value)
	}
}
