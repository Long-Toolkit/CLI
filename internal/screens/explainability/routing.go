package explainability

import (
	"strings"
)

func Heading(path []string) string {
	canonical, handled := CanonicalPath(path)
	if !handled || len(canonical) <= 1 {
		return "Explainability"
	}

	if tool, ok := ToolFromPath(canonical); ok {
		return tool.Name
	}

	parts := make([]string, 0, len(canonical))
	for _, part := range canonical {
		parts = append(parts, strings.TrimSpace(part))
	}
	return strings.Join(parts, " / ")
}

func CanonicalPath(path []string) ([]string, bool) {
	if len(path) == 0 {
		return nil, false
	}

	root := canonicalToken(path[0])
	if root != "explainability" {
		return nil, false
	}

	canonical := make([]string, 0, len(path))
	canonical = append(canonical, "explainability")

	for _, token := range path[1:] {
		tokenKey := canonicalToken(token)
		if toolID, ok := canonicalToolID(tokenKey); ok {
			canonical = append(canonical, toolID)
			continue
		}
		if listCommand, ok := canonicalToolListCommand(tokenKey); ok {
			canonical = append(canonical, listCommand)
			continue
		}
		canonical = append(canonical, token)
	}

	return canonical, true
}

func StripHelpFlags(path []string) ([]string, bool) {
	filtered := make([]string, 0, len(path))
	helpRequested := false
	for _, token := range path {
		switch strings.TrimSpace(token) {
		case "-h", "--help":
			helpRequested = true
		default:
			filtered = append(filtered, token)
		}
	}
	return filtered, helpRequested
}

func ContainsToken(path []string, token string) bool {
	for _, part := range path {
		if part == token {
			return true
		}
	}
	return false
}

func canonicalToken(raw string) string {
	return strings.TrimSpace(strings.ToLower(raw))
}

func canonicalToolID(token string) (string, bool) {
	for _, tool := range Tools() {
		if canonicalToken(tool.ID) == token {
			return tool.ID, true
		}
	}
	return "", false
}

func canonicalToolListCommand(token string) (string, bool) {
	for _, tool := range Tools() {
		if strings.TrimSpace(tool.ListCommand) == "" {
			continue
		}
		if canonicalToken(tool.ListCommand) == token {
			return tool.ListCommand, true
		}
	}
	return "", false
}
