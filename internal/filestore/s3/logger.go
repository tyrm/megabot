package s3

import "github.com/tyrm/megabot/internal/log"

type empty struct{}

var logger = log.WithPackageField(empty{})
