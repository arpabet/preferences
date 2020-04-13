// +build !ci

// +build linux,freebsd,openbsd,netbsd

package preferences

import (
	"os"
)

func AppDataDir(appName string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".config", groupName, appName)
}
