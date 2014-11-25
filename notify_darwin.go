// +build darwin

package notify

import "os/exec"

// Display shows the provided notification
func (n Notification) Display() error {
	return exec.Command("osascript", n.command()...).Run()
}
