package testdata

import (
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestFoo16(t *testing.T) {
	type args struct {
		in Bazzar
	}
	tests := []struct {
		name string
		args args
		want Bazzar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		if got := Foo16(tt.args.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo16() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
