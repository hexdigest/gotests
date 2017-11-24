package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestBook_Open(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(mc *minimock.Controller) *Book
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		b := tt.setup(mc)
		if err := b.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. Book.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_door_Open(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(mc *minimock.Controller) *door
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		d := tt.setup(mc)
		if err := d.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. door.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_xml_Open(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(mc *minimock.Controller) *xml
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		x := tt.setup(mc)
		if err := x.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. xml.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
