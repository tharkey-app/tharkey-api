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
	if v.Version != "(devel)" {
		t.Errorf("Expecting '%s' got '%s'", "(devel)", v.Version)
	}
}
