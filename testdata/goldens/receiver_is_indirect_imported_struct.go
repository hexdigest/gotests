package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func Test_someIndirectImportedStruct_Foo037(t *testing.T) {
	tests := []struct {
		name  string
		setup func(mc *minimock.Controller) *someIndirectImportedStruct
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		smtg := tt.setup(mc)
		smtg.Foo037()
	}
}
