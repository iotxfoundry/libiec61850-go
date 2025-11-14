package libiec61850go

import "testing"

func TestGetVersionString(t *testing.T) {
	version := GetVersionString()
	if version == "" {
		t.Errorf("GetVersionString() returned empty string")
	}
	t.Logf("Version: %s", version)
}
