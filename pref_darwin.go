// +build !ci
// +build !mobile,!ios

package preferences

import (
	"os"
	"path/filepath"
)

func AppDataDir(appName string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, "Library", "Preferences", GroupName, appName)
}

