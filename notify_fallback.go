// +build !darwin

package notify

import "errors"

// ErrorUnsupportedPlatform indicates that the current platform is not
// supported for error display. Because this package relies on Applescript
// (osascript) support, any platform other than Darwin (OS X) will fail with
// this error.
var ErrorUnsupportedPlatform = errors.New("unsupported platform")

// Show shows the provided notification
func Show(n commander) error {
	return UnsupportedPlatformError
}
