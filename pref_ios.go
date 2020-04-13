// +build !ci

// +build darwin,ios

package preferences

func AppDataDir(appName string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, "Library", "Preferences", groupName, appName)
}
