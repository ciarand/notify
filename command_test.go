package notify

import (
	"fmt"
	"strings"
	"testing"
)

var commandTests = []struct {
	desc         string
	notification Notification
	want         []string
}{
	{
		"basic notification",
		Notification{Text: "foo", Title: "bar"},
		[]string{"-e", `display notification "foo" with title "bar"`},
	},
	{
		"double quote semantics",
		Notification{Text: `foo"doublequote"`, Title: `bar"doublequote"`},
		[]string{"-e", `display notification "foo\"doublequote\"" with title "bar\"doublequote\""`},
	},
	{
		"single quote semantics",
		Notification{Text: `foo'singlequote'`, Title: `bar'singlequote'`},
		[]string{"-e", `display notification "foo'singlequote'" with title "bar'singlequote'"`},
	},
	{
		"with a subtitle",
		Notification{Text: `foo`, Title: `bar`, Subtitle: "baz"},
		[]string{
			"-e",
			`display notification "foo" with title "bar" subtitle "baz"`,
		},
	},
	{
		"with a subtitle containing quote marks",
		Notification{Text: `foo`, Title: `bar`, Subtitle: `baz"malicious"`},
		[]string{
			"-e",
			`display notification "foo" with title "bar" subtitle "baz\"malicious\""`,
		},
	},
	{
		"with a sound",
		Notification{Text: `foo`, Title: `bar`, Sound: "Funk"},
		[]string{
			"-e",
			`display notification "foo" with title "bar" sound name "Funk"`,
		},
	},
	{
		"with a sound contining quote marks",
		Notification{Text: `foo`, Title: `bar`, Sound: `Funk"malicious"`},
		[]string{
			"-e",
			`display notification "foo" with title "bar" sound name "Funk\"malicious\""`,
		},
	},
	{
		"with everything containing double quotes",
		Notification{Text: `"dq"`, Title: `"dq"`, Subtitle: `"dq"`, Sound: `"dq"`},
		[]string{
			"-e",
			`display notification "\"dq\"" with title "\"dq\"" subtitle "\"dq\"" sound name "\"dq\""`,
		},
	},
}

func TestCommandCorrectlyFormatsTheCommandString(t *testing.T) {
	t.Parallel()

	for i, perm := range commandTests {
		for j, s := range perm.notification.command() {
			if s != perm.want[j] {
				t.Errorf(`Error with %s (#%d).
    Got:      "%v"
    Expected: "%v"`, perm.desc, i+1, s, perm.want[j])
			}
		}
	}
}

type mockNotification struct {
	args []string
}

func (m mockNotification) command() []string {
	return m.args
}

func (m mockNotification) String() string {
	return fmt.Sprintf("osascript %s", strings.Join(m.command(), " "))
}

func TestShowReturnsAnErrorOnMalformedCommand(t *testing.T) {
	t.Parallel()

	m := mockNotification{[]string{"fail"}}

	if err := Show(m); err == nil {
		t.Errorf("Expected `%s` to fail, but it succeeded", m)
	}
}

func TestShowReturnsNilOnCorrectlyFormedCommand(t *testing.T) {
	t.Parallel()

	m := mockNotification{[]string{"-e", "random number"}}

	if err := Show(m); err != nil {
		t.Errorf("Expected `%s` to succeed, but it failed", m)
	}
}
