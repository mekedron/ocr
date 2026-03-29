package cli

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var errVersionShown = errors.New("version shown")

// NewRootCommand builds the command tree for ocr.
func NewRootCommand(version string) *cobra.Command {
	resolved := resolvedVersion(version)

	root := &cobra.Command{
		Use:           "ocr",
		Short:         "Capture a screenshot, OCR it, and print text to stdout.",
		SilenceErrors: true,
		SilenceUsage:  true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			showVersion, _ := cmd.Flags().GetBool("version")
			if showVersion {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), resolved)
				return errVersionShown
			}

			lang, _ := cmd.Flags().GetString("lang")
			silent, _ := cmd.Flags().GetBool("silent")
			return runCapture(cmd, lang, silent)
		},
	}

	root.Flags().BoolP("version", "v", false, "Show CLI version and exit.")
	root.Flags().StringP("lang", "l", "en-US", "OCR language(s), e.g. en-US+ru-RU. See: ocr languages.")
	root.Flags().BoolP("silent", "x", false, "Do not play sounds.")
	root.SetHelpCommand(&cobra.Command{Hidden: true})
	root.AddCommand(newLanguagesCommand())

	return root
}
