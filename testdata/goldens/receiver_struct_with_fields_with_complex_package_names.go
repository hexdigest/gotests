package testdata

import (
	"go/types"
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestImporter_Foo35(t *testing.T) {
	type args struct {
		t types.Type
	}
	tests := []struct {
		name  string
		setup func(mc *minimock.Controller) *Importer
		args  args
		want  *types.Var
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		i := tt.setup(mc)
		if got := i.Foo35(tt.args.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Importer.Foo35() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
