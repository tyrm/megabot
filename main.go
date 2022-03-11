package megabot

import "embed"

// Files contains static files required by the application
//go:embed active.*.toml
//go:embed web/static/css/*
//go:embed web/static/img/*
//go:embed web/static/js/*
//go:embed web/static/vendor/fontawesome-free-6.0.0-web/css/all.min.css
//go:embed web/static/vendor/fontawesome-free-6.0.0-web/webfonts/*
//go:embed web/static/vendor/fontawesome-free-6.0.0-web/LICENSE.txt
//go:embed web/template/*
var Files embed.FS
