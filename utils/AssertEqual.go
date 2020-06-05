package utils

import (
	"testing"
)

func AssertStringEqual(t *testing.T, target, have string) {
	if target != have {
		t.Errorf("\nWant got %s\nBut got %s", target, have)
	}
}

func AssertStringNotEqual(t *testing.T, target, have string) {
	if target == have {
		t.Errorf("\nGot same string%s\n", have)
	}
}

func AssertIntEqual(t *testing.T, target, have int) {
	if target != have {
		t.Errorf("\nWant got %d\nBut got %d", target, have)
	}
}
