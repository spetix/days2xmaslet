package out

import (
	"os"

	data "github.com/spetix/days2xmasleft/internal/data"
)

type baseOutput struct {
	Device *os.File
}

type BlockletOutput interface {
	Print(d data.Data)
}

func New(protocol string) BlockletOutput {

	switch protocol {
	case "json":
		return NewJsonOut(os.Stdout)
	case "raw":
		return NewRawOut(os.Stdout)
	case "waybar":
		return NewWaybarOut(os.Stdout)
	default:
		return NewRawOut(os.Stdout)
	}
}
