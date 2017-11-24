package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestFoo2(t *testing.T) {
	type args struct {
		in0 string
		in1 int
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
		Foo2(tt.args.in0, tt.args.in1)
	}
}
