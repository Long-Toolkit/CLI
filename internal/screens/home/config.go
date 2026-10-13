package home

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	toolkitconfig "long-toolkit-cli/config"
	"long-toolkit-cli/internal/shared/common"
	"long-toolkit-cli/internal/shared/model"
)

const (
	homeConfigEnvVar      = "LONG_TOOLKIT_HOME_CONFIG"
	defaultHomeConfigPath = "config/home.json"
)

type LabelledEntry struct {
	Label string `json:"label"`
	Text  string `json:"text"`
}

type Motivation struct {
	Quote       string `json:"quote"`
	Attribution string `json:"attribution"`
	ReadMore    string `json:"read_more"`
}

type TutorialStep struct {
	Title          string `json:"title"`
	Goal           string `json:"goal"`
	Outcome        string `json:"outcome"`
	Recommendation string `json:"recommendation"`
}

type Config struct {
	Actions    []model.CommandDef `json:"actions"`
	Authors    []LabelledEntry    `json:"authors"`
	Motivation Motivation         `json:"motivation"`
	Tutorial   []TutorialStep     `json:"tutorial_steps"`
}

var currentHomeConfig = defaultHomeConfig()

var homeActions = cloneCommandDefs(currentHomeConfig.Actions)

func ConfigureContent() error {
	config, err := loadHomeConfig()
	if err != nil {
		return err
	}
	currentHomeConfig = config
	homeActions = cloneCommandDefs(config.Actions)
	return nil
}

func Actions() []model.CommandDef {
	return cloneCommandDefs(homeActions)
}

func CurrentConfig() Config {
	return currentHomeConfig
}

func defaultHomeConfig() Config {
	return Config{
		Actions: []model.CommandDef{
			{
				Name:        "toolkit_tutorial",
				DisplayName: "Toolkit tutorial",
				Description: "Run the interactive toolkit tutorial.",
			},
			{
				Name:        "list_authors",
				DisplayName: "List authors",
				Description: "List the authors of the toolkit.",
			},
			{
				Name:        "motivation",
				DisplayName: "Motivation",
				Description: "Why we built Long-Toolkit.",
			},
		},
		Authors: []LabelledEntry{
			{
				Label: "1",
				Text: "Simon Provost, Postdoctoral Research @ McGill University (MEDomicsLab); " +
					"https://orcid.org/0000-0001-8402-5464",
			},
			{
				Label: "2",
				Text: "Bianca Branco (PhD student, University of Edinburgh; " +
					"https://orcid.org/0000-0003-0031-6200)",
			},
			{
				Label: "3",
				Text:  "Honghui Chen",
			},
			{
				Label: "4",
				Text:  "Isaac Tassella",
			},
			{
				Label: "5",
				Text: "Alex Kwong (Wellcome Senior Research Fellow, University of Edinburgh; " +
					"https://orcid.org/0000-0003-1953-2771)",
			},
		},
		Motivation: Motivation{
			Quote: "Longitudinal studies invest heavily in collecting and curating their data, yet machine " +
				"learning is either left aside or applied as if every wave were an unrelated feature. We " +
				"built Long-Toolkit so that experts of longitudinal studies can get their data ready, and " +
				"measure fairly whether conventional or longitudinal machine learning is worth it.",
			Attribution: "Simon Provost, Postdoctoral Research @ McGill University (MEDomicsLab) & Dr Alex Kwong, LEAP lab (University of Edinburgh)",
			ReadMore:    "https://long-toolkit.readthedocs.io/en/latest/getting-started/#motivation",
		},
		Tutorial: []TutorialStep{
			{
				Title: "Step 1: Stage A (Data Preparation)",
				Goal: "Use the data preparation tools to generate synthetic longitudinal data with controlled " +
					"temporal patterns, or to convert cohort files between CSV, Parquet and Stata formats.",
				Outcome: "Output: prepared longitudinal dataset(s) ready for preprocessing.",
				Recommendation: "When no built-in tool fits a dataset, write a custom script " +
					"and consider opening a pull request.",
			},
			{
				Title: "Step 2: Stage B (Data Preprocessing)",
				Goal: "Use the data preprocessing tools to clean, impute, harmonise and reshape longitudinal " +
					"datasets so they are ready for modelling.",
				Outcome: "Output: quality-checked dataset(s) ready for modelling.",
				Recommendation: "When no built-in tool fits a dataset, implement a custom workflow " +
					"and consider opening a pull request.",
			},
			{
				Title: "Step 3: Stage C (Machine Learning)",
				Goal: "Use the benchmarking tools to compare standard and longitudinal machine learning under " +
					"one shared protocol, and establish whether machine learning, and/or longitudinal machine " +
					"learning, is worth it on a dataset compared with a statistical baseline. Training " +
					"tools sit alongside them.",
				Outcome: "Output: saved models, benchmark rankings and evaluation summaries ready to explain or share.",
				Recommendation: "When no built-in tool fits a dataset, run a custom machine learning pipeline and " +
					"consider opening a pull request.",
			},
			{
				Title: "Step 4: Stage D (Explainability)",
				Goal: "Use the SHAP analysis tool to explain a saved model, e.g., the best estimator of a " +
					"benchmark, and see which features drive its predictions.",
				Outcome: "Output: SHAP values, feature-importance figures and report artefacts ready to share.",
				Recommendation: "When no built-in tool fits a model, run a custom explainability analysis and " +
					"consider opening a pull request.",
			},
		},
	}
}

func loadHomeConfig() (Config, error) {
	merged := defaultHomeConfig()

	bundledRaw, bundledSource, err := readBundledHomeConfig()
	if err != nil {
		return merged, err
	}
	if len(bundledRaw) > 0 {
		bundledConfig, err := decodeHomeConfig(bundledRaw, bundledSource)
		if err != nil {
			return merged, err
		}
		merged = mergeHomeConfig(merged, bundledConfig)
	}

	overridePath, hasOverride, err := findHomeOverrideConfigPath()
	if err != nil {
		return merged, err
	}
	if !hasOverride {
		if len(merged.Actions) == 0 {
			return merged, errors.New("home config produced zero actions; keep at least one action")
		}
		return merged, nil
	}

	overrideRaw, err := os.ReadFile(overridePath)
	if err != nil {
		return merged, fmt.Errorf("failed to read home config at %s: %w", overridePath, err)
	}

	overrideConfig, err := decodeHomeConfig(overrideRaw, overridePath)
	if err != nil {
		return merged, err
	}
	merged = mergeHomeConfig(merged, overrideConfig)

	if len(merged.Actions) == 0 {
		return merged, errors.New("home config produced zero actions; keep at least one action")
	}
	return merged, nil
}

func readBundledHomeConfig() ([]byte, string, error) {
	raw, embedErr := toolkitconfig.ReadBundledConfigFile(defaultHomeConfigPath)
	if embedErr == nil {
		return raw, defaultHomeConfigPath, nil
	}

	configPath, hasConfig, pathErr := findBundledHomeConfigPath()
	if pathErr != nil {
		return nil, "", fmt.Errorf(
			"failed to load bundled home config, embedded error: %v, filesystem error: %w",
			embedErr,
			pathErr,
		)
	}
	if !hasConfig {
		return nil, "", nil
	}

	raw, err := os.ReadFile(configPath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read home config at %s: %w", configPath, err)
	}
	return raw, configPath, nil
}

func decodeHomeConfig(raw []byte, source string) (Config, error) {
	var decoded Config
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return Config{}, fmt.Errorf("invalid JSON in %s: %w", source, err)
	}
	return decoded, nil
}

func mergeHomeConfig(base Config, override Config) Config {
	merged := base

	if len(override.Actions) > 0 {
		merged.Actions = sanitiseActions(override.Actions)
	}
	if len(override.Authors) > 0 {
		merged.Authors = sanitiseLabelledEntries(override.Authors)
	}
	if motivation := sanitiseMotivation(override.Motivation); motivation.Quote != "" {
		merged.Motivation = motivation
	}
	if len(override.Tutorial) > 0 {
		merged.Tutorial = sanitiseTutorialSteps(override.Tutorial)
	}
	return merged
}

func findHomeOverrideConfigPath() (string, bool, error) {
	if value := strings.TrimSpace(os.Getenv(homeConfigEnvVar)); value != "" {
		resolved, err := common.ResolveConfigPath(value)
		if err != nil {
			return "", false, err
		}
		return resolved, true, nil
	}
	return "", false, nil
}

func findBundledHomeConfigPath() (string, bool, error) {
	candidates := []string{
		defaultHomeConfigPath,
		filepath.Join("cli", defaultHomeConfigPath),
		filepath.Join("src", "cli", defaultHomeConfigPath),
	}

	if root, err := common.FindProjectRoot(); err == nil {
		candidates = append(
			candidates,
			filepath.Join(root, "src", "cli", defaultHomeConfigPath),
			filepath.Join(root, "cli", defaultHomeConfigPath),
		)
	}

	for _, candidate := range candidates {
		resolved, err := common.ResolveConfigPath(candidate)
		if err != nil {
			return "", false, err
		}
		_, statErr := os.Stat(resolved)
		if statErr == nil {
			return resolved, true, nil
		}
		if !os.IsNotExist(statErr) {
			return "", false, statErr
		}
	}

	return "", false, nil
}

func sanitiseActions(actions []model.CommandDef) []model.CommandDef {
	cleaned := make([]model.CommandDef, 0, len(actions))
	for _, action := range actions {
		name := strings.TrimSpace(action.Name)
		displayName := strings.TrimSpace(action.DisplayName)
		desc := strings.TrimSpace(action.Description)
		if name == "" {
			continue
		}
		if displayName == "" {
			displayName = "Unnamed action"
		}
		if desc == "" {
			desc = displayName
		}
		cleaned = append(cleaned, model.CommandDef{
			Name:        name,
			DisplayName: displayName,
			Description: desc,
		})
	}
	return cleaned
}

func sanitiseLabelledEntries(entries []LabelledEntry) []LabelledEntry {
	cleaned := make([]LabelledEntry, 0, len(entries))
	for index, entry := range entries {
		text := strings.TrimSpace(entry.Text)
		if text == "" {
			continue
		}
		label := strings.TrimSpace(entry.Label)
		if label == "" {
			label = strconv.Itoa(index + 1)
		}
		cleaned = append(cleaned, LabelledEntry{
			Label: label,
			Text:  text,
		})
	}
	return cleaned
}

func sanitiseMotivation(motivation Motivation) Motivation {
	return Motivation{
		Quote:       strings.TrimSpace(motivation.Quote),
		Attribution: strings.TrimSpace(motivation.Attribution),
		ReadMore:    strings.TrimSpace(motivation.ReadMore),
	}
}

func sanitiseTutorialSteps(steps []TutorialStep) []TutorialStep {
	cleaned := make([]TutorialStep, 0, len(steps))
	for _, step := range steps {
		title := strings.TrimSpace(step.Title)
		goal := strings.TrimSpace(step.Goal)
		outcome := strings.TrimSpace(step.Outcome)
		recommendation := strings.TrimSpace(step.Recommendation)
		if title == "" || goal == "" || outcome == "" {
			continue
		}
		cleaned = append(cleaned, TutorialStep{
			Title:          title,
			Goal:           goal,
			Outcome:        outcome,
			Recommendation: recommendation,
		})
	}
	return cleaned
}

func cloneCommandDefs(input []model.CommandDef) []model.CommandDef {
	output := make([]model.CommandDef, len(input))
	copy(output, input)
	return output
}
