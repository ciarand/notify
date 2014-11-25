// +build !darwin

package notify

// Show shows the provided notification
func Show(n commander) error {
	return ErrorUnsupportedPlatform
}
