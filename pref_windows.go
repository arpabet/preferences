// +build !ci

// +build !mobile,!android,!ios

package preferences

import (
	"os"
	"path/filepath"
)

func AppDataDir(appName string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, "AppData", "Roaming", groupName, appName)
}

