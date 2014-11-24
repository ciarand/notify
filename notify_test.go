package notify

import "testing"

var notifyTests = []struct {
	desc string
	got  Notification
	want Notification
}{
	{
		`NewNotification("foo", "bar")`,
		NewNotification("foo", "bar"),
		Notification{Title: "foo", Text: "bar"},
	},
	{
		`NewSubtitledNotification("foo", "bar", "baz")`,
		NewSubtitledNotification("foo", "bar", "baz"),
		Notification{Title: "foo", Subtitle: "bar", Text: "baz"},
	},
	{
		`NewNotificationWithSound("foo", "bar", "baz")`,
		NewNotificationWithSound("foo", "bar", "baz"),
		Notification{Title: "foo", Text: "bar", Sound: "baz"},
	},
	{
		`NewSubtitledNotificationWithSound("foo", "bar", "baz", "qux")`,
		NewSubtitledNotificationWithSound("foo", "bar", "baz", "qux"),
		Notification{Title: "foo", Subtitle: "bar", Text: "baz", Sound: "qux"},
	},
}

func TestNewNotificationHelpers(t *testing.T) {
	t.Parallel()

	for _, perm := range notifyTests {
		assertEquals(t, perm.got, perm.want)
	}
}
