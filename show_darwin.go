// +build darwin

package notify

import (
	"fmt"
	"strconv"
)

const showCommand = "osascript"

// command generates the command to use.
func (n Notification) command() []string {
	cmd := fmt.Sprintf("display notification %s with title %s",
		strconv.Quote(n.Text), strconv.Quote(n.Title))

	if n.Subtitle != "" {
		cmd = fmt.Sprintf("%s subtitle %s", cmd, strconv.Quote(n.Subtitle))
	}

	if n.Sound != "" {
		cmd = fmt.Sprintf("%s sound name %s", cmd, strconv.Quote(n.Sound))
	}

	return []string{"-e", cmd}
}
