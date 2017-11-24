package bar

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestBar_Bar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		setup   func(mc *minimock.Controller) *Bar
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		b := tt.setup(mc)
		if err := b.Bar(tt.args.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.Bar() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
