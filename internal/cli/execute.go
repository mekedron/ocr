package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Execute runs the command tree and returns an exit code.
func Execute(ctx context.Context, args []string, version string, out io.Writer, errOut io.Writer) int {
	root := NewRootCommand(version)
	root.SetArgs(args)
	root.SetOut(out)
	root.SetErr(errOut)

	err := root.ExecuteContext(ctx)
	if err == nil {
		return 0
	}
	if errors.Is(err, errVersionShown) {
		return 0
	}
	if errors.Is(err, errCancelled) {
		return 0
	}

	message := strings.TrimSpace(err.Error())
	if message != "" {
		_, _ = fmt.Fprintln(errOut, message)
	}
	return 1
}
