package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func Test_initFuncStruct_init(t *testing.T) {
	tests := []struct {
		name  string
		setup func(mc *minimock.Controller) initFuncStruct
		want  int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		i := tt.setup(mc)
		if got := i.init(); got != tt.want {
			t.Errorf("%q. initFuncStruct.init() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_initFieldStruct_getInit(t *testing.T) {
	tests := []struct {
		name  string
		setup func(mc *minimock.Controller) initFieldStruct
		want  int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		i := tt.setup(mc)
		if got := i.getInit(); got != tt.want {
			t.Errorf("%q. initFieldStruct.getInit() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
