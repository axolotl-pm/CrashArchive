package crashreport

import "testing"

func TestVersionStringRoundTrip(t *testing.T) {
	testVersions := [...]string{"1.6.2", "1.6.2-1800", "4.0.0-ALPHA1"}

	for _, versionText := range testVersions {
		version, err := NewVersionString(versionText, 0, false)
		if err != nil {
			t.Fatalf("failed to parse %s: %v", versionText, err)
		}
		if version.Get(true) != versionText {
			t.Fatalf("Bad version string, expected %s, got %s", versionText, version.Get(true))
		}
	}
}

func TestVersionStringDevelopmentBuild(t *testing.T) {
	version, err := NewVersionString("1.6.2", 178, true)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}
	if got := version.Get(true); got != "1.6.2+dev.178" {
		t.Fatalf("expected 1.6.2+dev.178, got %s", got)
	}
	if got := version.Get(false); got != "1.6.2+dev" {
		t.Fatalf("expected 1.6.2+dev, got %s", got)
	}
}

func TestVersionStringRejectsIncomplete(t *testing.T) {
	for _, versionText := range [...]string{"1.6", "1.6dev", "1.6.2dev"} {
		if _, err := NewVersionString(versionText, 0, false); err == nil {
			t.Fatalf("expected an error for %s", versionText)
		}
	}
}
