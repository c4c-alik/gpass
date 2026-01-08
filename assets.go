package main

import (
	"embed"
	"io/fs"
)

//go:embed assets/*
var assets embed.FS

// Assets returns the embedded assets
func Assets() fs.FS {
	return assets
}
