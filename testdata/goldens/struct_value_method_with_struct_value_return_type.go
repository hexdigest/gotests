package testdata

import (
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestBar_Foo9(t *testing.T) {
	tests := []struct {
		name  string
		setup func(mc *minimock.Controller) Bar
		want  Bar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		b := tt.setup(mc)
		if got := b.Foo9(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Bar.Foo9() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
