package megabot

import "embed"

// Files contains static files required by the application
//go:embed active.*.toml
//go:embed web/static/*  web/template/*
var Files embed.FS
