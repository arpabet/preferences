// +build !ci

// +build linux,freebsd,openbsd,netbsd

package preferences

import (
	"os"
	"path/filepath"
)

func AppDataDir(appName string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".config", groupName, appName)
}
