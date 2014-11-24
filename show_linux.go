// +build linux

package notify

const showCommand = "notify-send"

// command generates the command to use.
func (n Notification) command() []string {
	return []string{n.Title, n.Text}
}
