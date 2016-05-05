// +build !windows

package service

// ProcessWindowsControlEvents is not used on non-windows platforms.
func ProcessWindowsControlEvents(stopCallback func()) {
}

// InstallWindowsService is not used on non-windows platforms.
func InstallWindowsService() error {
	return nil
}

// UninstallWindowsService is not used on non-windows platforms.
func UninstallWindowsService(name string) error {
	return nil
}
