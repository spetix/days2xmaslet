package blockletapi

import (
	"os"
	"time"

	"github.com/spetix/days2xmasleft/internal/dateutil"
)

type JsonOut struct {
	baseOutput
}

func NewJsonOut(device *os.File) BlockletOutput {
	return &JsonOut{
		baseOutput: baseOutput{
			Device: device,
		},
	}
}

func (j *JsonOut) Print(unit time.Duration, rerenderOptions *RenderOptions) {

	j.Device.Write([]byte(rerenderOptions.Label))
	j.Device.Write([]byte(dateutil.Format(unit, rerenderOptions.Unit)))
}
