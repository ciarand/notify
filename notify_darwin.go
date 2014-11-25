// +build darwin

package notify

import "os/exec"

// Show shows the provided notification
func Show(n commander) error {
	return exec.Command(showCommand, n.command()...).Run()
}
