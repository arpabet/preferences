// +build !ci

// +build android mobile

package preferences

import "path/filepath"

func AppDataDir(appName string) string {
	return filepath.Join("/data", "data", groupName, appName)
}

