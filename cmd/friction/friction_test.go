package main

import (
	"friction/loggers"
	"testing"
)

func TestIsZapTrue(t *testing.T) {
	var logger loggers.Logger
	logger = &loggers.Zap{}

	got := isZap(logger)
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
