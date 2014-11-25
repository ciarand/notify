package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ciarand/notify"
)

var (
	text     *string
	title    *string
	subtitle *string
	sound    *string
)

func init() {
	text = flag.String("text", "", "The required text for the notification")
	title = flag.String("title", "", "The required title of the notification")
	subtitle = flag.String("subtitle", "", "An optional subtitle for the notification")
	sound = flag.String("sound", "", "An optional sound name to play")
}

func main() {
	flag.Parse()
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	if *title == "" {
		dief("Title cannot be blank")
	}

	if *text == "" {
		dief("Text cannot be blank")
	}

	n := notify.NewSubtitledNotificationWithSound(*title, *subtitle, *text, *sound)

	if err := n.Display(); err != nil {
		dief("Error showing notification: %s", err)
	}
}

func dief(s string, args ...interface{}) {
	fmt.Printf(s+"\n", args...)

	os.Exit(1)
}
