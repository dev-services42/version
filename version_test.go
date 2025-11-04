//nolint:testpackage
package version

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	t.Parallel()

	t.Run("returns version from build info", func(t *testing.T) {
		t.Parallel()

		res := GetVersion()
		for _, ver := range []string{"(devel)", "unknown-local"} {
			if ver == res {
				return
			}
		}

		t.Errorf("unexpected version %q", res)
	})

	t.Run("returns explicitly set version when version variable is set", func(t *testing.T) {
		t.Parallel()

		// Save the original value
		original := version
		defer func() { version = original }()

		// Set explicit version
		version = "v1.2.3"
		if GetVersion() != "v1.2.3" {
			t.Error("version should be set to v1.2.3")
		}
	})
}
