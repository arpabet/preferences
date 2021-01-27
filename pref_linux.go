// +build linux openbsd freebsd netbsd
// +build !android

package preferences

import (
	"os"
	"path/filepath"
)

func AppDataDir(appName string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".config", GroupName, appName)
}
