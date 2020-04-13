// +build !ci

// +build !mobile,!ios

package preferences


func AppDataDir(appName string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join("Library", "Preferences", groupName, appName)
}

