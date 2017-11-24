package testdata

import (
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestFoo10(t *testing.T) {
	type args struct {
		m map[string]int32
	}
	tests := []struct {
		name string
		args args
		want map[string]*Bar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		if got := Foo10(tt.args.m); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo10() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
