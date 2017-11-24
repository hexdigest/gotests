package testdata

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestFoo18(t *testing.T) {
	type args struct {
		t *os.File
	}
	tests := []struct {
		name string
		args args
		want *os.File
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		if got := Foo18(tt.args.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo18() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
