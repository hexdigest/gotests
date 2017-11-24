package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func Test_name_Name(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		n    name
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		if got := tt.n.Name(tt.args.n); got != tt.want {
			t.Errorf("%q. name.Name() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestName_Name1(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name  string
		setup func(mc *minimock.Controller) *Name
		args  args
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		n := tt.setup(mc)
		if got := n.Name1(tt.args.n); got != tt.want {
			t.Errorf("%q. Name.Name1() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestName_Name2(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		setup func(mc *minimock.Controller) *Name
		args  args
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		n := tt.setup(mc)
		if got := n.Name2(tt.args.name); got != tt.want {
			t.Errorf("%q. Name.Name2() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestName_Name3(t *testing.T) {
	type args struct {
		nn string
	}
	tests := []struct {
		name     string
		setup    func(mc *minimock.Controller) *Name
		args     args
		wantName string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		n := tt.setup(mc)
		if gotName := n.Name3(tt.args.nn); gotName != tt.wantName {
			t.Errorf("%q. Name.Name3() = %v, want %v", tt.name, gotName, tt.wantName)
		}
	}
}
