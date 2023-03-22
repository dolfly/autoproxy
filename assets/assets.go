package assets

import "embed"

//go:embed *.json *.jpg *.mmdb html/*
var FS embed.FS
