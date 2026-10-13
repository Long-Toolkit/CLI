package tools

import (
	"errors"

	"github.com/charmbracelet/huh"

	"long-toolkit-cli/internal/shared/components"
)

func IsFlowCancelled(err error) bool {
	return errors.Is(err, huh.ErrUserAborted) || errors.Is(err, components.ErrPathPickerCancelled)
}
