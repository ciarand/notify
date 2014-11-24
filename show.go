package notify

import "os/exec"

type commander interface {
	command() []string
}

// Show shows the provided notification
func Show(n commander) error {
	return exec.Command(showCommand, n.command()...).Run()
}
