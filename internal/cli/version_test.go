package cli

import (
	"runtime/debug"
	"testing"
)

func TestResolvedVersion_ExplicitVersion(t *testing.T) {
	got := resolvedVersion("v1.2.3")
	if got != "v1.2.3" {
		t.Errorf("expected v1.2.3, got %s", got)
	}
}

func TestResolvedVersion_DevFallsBackToBuildInfo(t *testing.T) {
	original := readBuildInfo
	defer func() { readBuildInfo = original }()

	readBuildInfo = func() (*debug.BuildInfo, bool) {
		return &debug.BuildInfo{
			Settings: []debug.BuildSetting{
				{Key: vcsRevisionKey, Value: "abcdef123456789"},
			},
		}, true
	}

	got := resolvedVersion("dev")
	if got != "abcdef123456" {
		t.Errorf("expected abcdef123456, got %s", got)
	}
}

func TestResolvedVersion_DirtyRevision(t *testing.T) {
	original := readBuildInfo
	defer func() { readBuildInfo = original }()

	readBuildInfo = func() (*debug.BuildInfo, bool) {
		return &debug.BuildInfo{
			Settings: []debug.BuildSetting{
				{Key: vcsRevisionKey, Value: "abcdef123456789"},
				{Key: vcsModifiedKey, Value: "true"},
			},
		}, true
	}

	got := resolvedVersion("")
	if got != "abcdef123456-dirty" {
		t.Errorf("expected abcdef123456-dirty, got %s", got)
	}
}

func TestResolvedVersion_NoBuildInfoReturnsDev(t *testing.T) {
	original := readBuildInfo
	defer func() { readBuildInfo = original }()

	readBuildInfo = func() (*debug.BuildInfo, bool) {
		return nil, false
	}

	got := resolvedVersion("")
	if got != devVersion {
		t.Errorf("expected %s, got %s", devVersion, got)
	}
}
