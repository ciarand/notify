// +build !linux,!darwin

package notify

import "fmt"

const showCommand = "echo"

func (n Notification) command() []string {
	return []string{fmt.Sprintf("%s: %s", n.Title, n.Text)}
}
