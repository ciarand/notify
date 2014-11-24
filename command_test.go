package notify

import (
	"fmt"
	"strings"
	"testing"
)

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
	return fmt.Sprintf("%s %s", showCommand, strings.Join(m.command(), " "))
}

func TestShowReturnsAnErrorOnMalformedCommand(t *testing.T) {
	t.Parallel()

	if failingMock.command() == nil {
		t.Skipf("no failingMock available for this platform")
		return
	}

	if err := Show(failingMock); err == nil {
		t.Errorf("Expected `%s` to fail, but it succeeded", failingMock)
	}
}

func TestShowReturnsNilOnCorrectlyFormedCommand(t *testing.T) {
	t.Parallel()

	if succeedingMock.command() == nil {
		t.Skipf("no succeedingMock available for this platform")
		return
	}

	if err := Show(succeedingMock); err != nil {
		t.Errorf("Expected `%s` to succeed, but it failed", succeedingMock)
	}
}
