package main

import (
	"context"
	"testing"
)

func TestGetVersion(t *testing.T) {

	v, err := getVersion(context.Background())
	if err != nil {
		t.Error(err)
	}
	if v == nil {
		t.Error("missing version")
	}
}
