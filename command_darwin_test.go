// +build darwin

package notify

import (
	"fmt"
	"strings"
	"testing"
)

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

	if err := Show(failingMock); err == nil {
		t.Errorf("Expected `%s` to fail, but it succeeded", failingMock)
	}
}

func TestShowReturnsNilOnCorrectlyFormedCommand(t *testing.T) {
	t.Parallel()

	if err := Show(succeedingMock); err != nil {
		t.Errorf("Expected `%s` to succeed, but it failed", succeedingMock)
	}
}
