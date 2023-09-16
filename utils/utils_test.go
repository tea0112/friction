package utils

import (
	"errors"
	"strconv"
	"testing"
)

func TestGetIdFromPath(t *testing.T) {
	var want int64 = 9223372036854775807
	got, err := GetIdFromPath("/9223372036854775807")

	if err == nil && want != got {
		t.Errorf("want %q, got %q", want, got)
	}

	_, err = GetIdFromPath("/ab1")

	if errors.Is(err, strconv.ErrSyntax) == false {
		t.Errorf("got %q, want %q", err, strconv.ErrSyntax)
	}
}