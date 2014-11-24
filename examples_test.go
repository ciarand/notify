package notify_test

import "github.com/ciarand/notify"

func ExampleNewNotification() {
	notify.Show(notify.NewNotification("Title", "Text"))
}
