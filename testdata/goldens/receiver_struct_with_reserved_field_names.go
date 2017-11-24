package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestReserved_DontFail(t *testing.T) {
	tests := []struct {
		name  string
		setup func(mc *minimock.Controller) *Reserved
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		r := tt.setup(mc)
		if got := r.DontFail(); got != tt.want {
			t.Errorf("%q. Reserved.DontFail() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
