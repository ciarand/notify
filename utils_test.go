package notify

import (
	"reflect"
	"testing"
)

func assertEquals(t *testing.T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, expected %v", got, want)
	}
}
