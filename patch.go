package main

import (
	"fmt"
	"runtime/debug"

	"github.com/goreleaser/goreleaser/pkg/config"
)

type Project struct {
	config.Project

	Archive config.Archive
}

const none = "none"

//nolint:gochecknoinits
func init() {
	if info, available := debug.ReadBuildInfo(); available && commit == none {
		version = info.Main.Version
		commit = fmt.Sprintf("%s, mod sum: %s", commit, info.Main.Sum)
	}
}
