package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestDoctor_SayHello(t *testing.T) {
	type args struct {
		r *Person
	}
	tests := []struct {
		name  string
		setup func(mc *minimock.Controller) *Doctor
		args  args
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		d := tt.setup(mc)
		if got := d.SayHello(tt.args.r); got != tt.want {
			t.Errorf("%q. Doctor.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
