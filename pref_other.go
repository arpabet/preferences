// +build ci !linux,!darwin,!windows,!freebsd,!openbsd,!netbsd

package preferences

func AppDataDir(appName string) string {
	return filepath.Join("/tmp", GroupName, appName)
}
