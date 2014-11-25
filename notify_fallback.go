// +build !darwin

package notify

// Display shows the provided notification
func (n Notification) Display() error {
	return ErrorUnsupportedPlatform
}
