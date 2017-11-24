package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestFoo1(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		Foo1()
	}
}
