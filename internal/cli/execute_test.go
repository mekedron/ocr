package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestExecute_ShowsVersion(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	code := Execute(context.Background(), []string{"--version"}, "v1.0.0", &out, &errOut)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	if got := strings.TrimSpace(out.String()); got != "v1.0.0" {
		t.Errorf("expected v1.0.0, got %s", got)
	}
}

func TestExecute_ShowsHelp(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	code := Execute(context.Background(), []string{"--help"}, "v1.0.0", &out, &errOut)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(out.String(), "ocr") {
		t.Errorf("expected help to mention ocr, got %s", out.String())
	}
}

func TestExecute_UnknownFlagFails(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	code := Execute(context.Background(), []string{"--unknown-flag"}, "v1.0.0", &out, &errOut)
	if code != 1 {
		t.Errorf("expected exit code 1, got %d", code)
	}
}
