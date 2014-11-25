// +build !darwin

package notify_test

import (
	"fmt"

	"github.com/ciarand/notify"
)

func ExampleErrorUnsupportedPlatformCheck() {
	n := notify.NewNotification("Error", "Foobar failed, the defenses are failing!")

	if err := n.Display(); err == notify.ErrorUnsupportedPlatform {
		fmt.Println("You're not using Darwin!")
	} else if err != nil {
		fmt.Println("An error happened when running the command!")
	} else {
		fmt.Println("Everything went great!")
	}
	// Output:
	// You're not using Darwin!
}
