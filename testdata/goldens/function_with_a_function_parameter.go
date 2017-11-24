package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestFoo13(t *testing.T) {
	type args struct {
		f func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		if err := Foo13(tt.args.f); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo13() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
