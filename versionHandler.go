package main

import (
	"context"
	"fmt"
	"runtime/debug"
)

var versionHandler = jsonHandler[VersionResponse](getVersion)

type VersionResponse struct {
	Version string `json:"version"`
}

func getVersion(ctx context.Context) (VersionResponse, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return VersionResponse{}, fmt.Errorf("undefined version")
	}
	return VersionResponse{
		Version: info.Main.Version,
	}, nil
}
