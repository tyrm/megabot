package megabot

import "embed"

// Files contains static files required by the application
//go:embed locales/active.*.toml
//go:embed test/fixture.yml
//go:embed web/static/css/bootstrap.min.css
//go:embed web/static/css/bootstrap.min.css.map
//go:embed web/static/css/error.min.css
//go:embed web/static/css/login.min.css
//go:embed web/static/img/*
//go:embed web/static/js/bootstrap.bundle.min.js
//go:embed web/static/js/bootstrap.bundle.min.js.map
//go:embed web/static/vendor/fontawesome-free-6.0.0-web/css/all.min.css
//go:embed web/static/vendor/fontawesome-free-6.0.0-web/webfonts/*
//go:embed web/static/vendor/fontawesome-free-6.0.0-web/LICENSE.txt
//go:embed web/template/*
var Files embed.FS
