package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestPerson_SayHello(t *testing.T) {
	type args struct {
		r *Person
	}
	tests := []struct {
		name  string
		setup func(mc *minimock.Controller) *Person
		args  args
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		p := tt.setup(mc)
		if got := p.SayHello(tt.args.r); got != tt.want {
			t.Errorf("%q. Person.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
