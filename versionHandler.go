package main

import (
	"context"
	"fmt"
	"runtime/debug"
)

var versionHandler = jsonHandler(getVersion)

func getVersion(ctx context.Context) (interface{}, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, fmt.Errorf("undefined version")
	}
	return struct {
		Version string `json:"version"`
	}{
		Version: info.Main.Version,
	}, nil
}
