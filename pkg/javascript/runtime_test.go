package javascript

import (
	"testing"
)

func TestNewRuntime(t *testing.T) {
	r := NewRuntime()
	if err := r.VM.Set("my.val", 1); err != nil {
		t.Fatalf("Failed to set myval=%d: %s", 1, err)
	}

	v, err := r.VM.Get("my.val")
	if err != nil {
		t.Fatalf("failed to get myval: %s", err)
	}
	iv, err := v.ToInteger()
	if err != nil {
		t.Error(err)
	}

	if iv != 1 {
		t.Errorf("Expected %d, got %d", 1, iv)
	}
}
