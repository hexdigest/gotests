package testdata

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestSameName(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		got, err := SameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. SameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_sameName(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		got, err := sameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. sameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. sameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSameTypeName_SameName(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(mc *minimock.Controller) *SameTypeName
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		t := tt.setup(mc)
		got, err := t.SameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SameTypeName.SameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. SameTypeName.SameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSameTypeName_sameName(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(mc *minimock.Controller) *SameTypeName
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		t := tt.setup(mc)
		got, err := t.sameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SameTypeName.sameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. SameTypeName.sameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_sameTypeName_SameName(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(mc *minimock.Controller) *sameTypeName
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		t := tt.setup(mc)
		got, err := t.SameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. sameTypeName.SameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. sameTypeName.SameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_sameTypeName_sameName(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(mc *minimock.Controller) *sameTypeName
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		t := tt.setup(mc)
		got, err := t.sameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. sameTypeName.sameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. sameTypeName.sameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
