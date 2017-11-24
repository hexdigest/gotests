package foo

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestFoo_Foo(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		setup   func(mc *minimock.Controller) *Foo
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		f := tt.setup(mc)
		if err := f.Foo(tt.args.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo.Foo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
