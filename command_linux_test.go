// +build linux

package notify

var failingMock = mockNotification{[]string{""}}
var succeedingMock = mockNotification{[]string{"title", "body text"}}

var commandTests = []struct {
	desc         string
	notification Notification
	want         []string
}{
	{
		"basic notification",
		Notification{Title: "foo", Text: "bar"},
		[]string{"foo", "bar"},
	},
	{
		"double quote semantics",
		Notification{Title: `foo"doublequote"`, Text: `bar"doublequote"`},
		[]string{`foo"doublequote"`, `bar"doublequote"`},
	},
	{
		"single quote semantics",
		Notification{Title: `foo'singlequote'`, Text: `bar'singlequote'`},
		[]string{`foo'singlequote'`, `bar'singlequote'`},
	},
}
