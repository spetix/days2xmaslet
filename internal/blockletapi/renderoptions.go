package blockletapi

import (
	"time"
)

type RenderOptions struct {
	Label           string
	Format          string
	ForegroundColor string
	BackgroundColor string
	Unit            time.Duration
}
