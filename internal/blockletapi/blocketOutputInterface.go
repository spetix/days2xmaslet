package blockletapi

import (
	"os"
	"time"
)

type baseOutput struct {
	Device *os.File
}

type BlockletOutput interface {
	Print(timeToXMas time.Duration, rednderOptions *RenderOptions)
}
