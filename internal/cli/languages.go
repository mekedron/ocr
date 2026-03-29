package cli

import (
	"fmt"

	"github.com/mekedron/ocr/internal/vision"
	"github.com/spf13/cobra"
)

func newLanguagesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "languages",
		Short: "List supported OCR languages.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			langs := vision.SupportedLanguages()
			if langs == "" {
				return fmt.Errorf("failed to query supported languages")
			}
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), langs)
			return nil
		},
	}
}
