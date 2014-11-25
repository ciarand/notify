package notify

import (
	"errors"
	"fmt"
	"strconv"
)

// ErrorUnsupportedPlatform indicates that the current platform is not
// supported for error display. Because this package relies on Applescript
// (osascript) support, any platform other than Darwin (OS X) will fail with
// this error.
var ErrorUnsupportedPlatform = errors.New("unsupported platform")

// Notification represents a notification that will be displayed.
type Notification struct {
	// The (required) title for the notification.
	Title string
	// An (optional) subtitle for the notification.
	Subtitle string
	// The (required) text to display in the notification.
	Text string
	// An (optional) path to a sound to play. Possible choices are pulled from
	// `~/Library/Sounds` and `/System/Library/Sounds`.
	Sound string
}

// NewNotification creates a new notification with the provided title and text.
func NewNotification(title, text string) Notification {
	return Notification{
		Title: title,
		Text:  text,
	}
}

// NewSubtitledNotification creates a new notification with the provided title,
// subtitle, and text.
func NewSubtitledNotification(title, sub, text string) Notification {
	n := NewNotification(title, text)

	n.Subtitle = sub

	return n
}

// NewNotificationWithSound creates a new notification with the provided title,
// text, and sound.
func NewNotificationWithSound(title, text, sound string) Notification {
	n := NewNotification(title, text)

	n.Sound = sound

	return n
}

// NewSubtitledNotificationWithSound creates a new notification with the
// provided title, subtitle, text, and sound.
func NewSubtitledNotificationWithSound(title, sub, text, sound string) Notification {
	n := NewNotification(title, text)

	n.Subtitle = sub
	n.Sound = sound

	return n
}

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
