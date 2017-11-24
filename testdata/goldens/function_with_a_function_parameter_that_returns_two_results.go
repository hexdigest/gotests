package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestFoo15(t *testing.T) {
	type args struct {
		f func(string) (string, error)
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
		if err := Foo15(tt.args.f); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo15() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
