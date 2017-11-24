package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestFoo3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		Foo3(tt.args.s)
	}
}
