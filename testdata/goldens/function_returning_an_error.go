package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestFoo5(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		got, err := Foo5()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo5() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Foo5() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
