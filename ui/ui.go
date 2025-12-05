package ui

import "embed"

//go:embed dist/*
//go:embed index.html
//go:embed index.css
var Content embed.FS
