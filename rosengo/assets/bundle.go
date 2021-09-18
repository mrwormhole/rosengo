package assets

import "embed"

//go:embed images/* sounds/*
var Bundle embed.FS
