package cli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mekedron/ocr/internal/vision"
	"github.com/spf13/cobra"
)

var errCancelled = errors.New("screenshot cancelled")

func runCapture(cmd *cobra.Command, lang string, silent bool) error {
	tmpFile, err := os.CreateTemp("", "ocr-*.png")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	_ = tmpFile.Close()
	defer func() { _ = os.Remove(tmpPath) }()

	captureArgs := []string{"-i", tmpPath}
	if silent {
		captureArgs = []string{"-ix", tmpPath}
	}
	captureCmd := exec.CommandContext(cmd.Context(), "screencapture", captureArgs...)
	if err := captureCmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() == 1 {
			return errCancelled
		}
		return fmt.Errorf("screencapture failed: %w", err)
	}

	info, err := os.Stat(tmpPath)
	if err != nil || info.Size() == 0 {
		return errCancelled
	}

	languages := strings.Split(lang, "+")
	text := vision.RecognizeText(tmpPath, languages)
	if text == "" {
		return nil
	}

	_, _ = fmt.Fprint(cmd.OutOrStdout(), text)
	return nil
}
