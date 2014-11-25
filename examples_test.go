package notify_test

import "github.com/ciarand/notify"

func ExampleNewNotification() {
	notify.NewNotification("Title", "Text").Display()
}
