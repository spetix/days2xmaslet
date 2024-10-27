package blockletapi

import (
	"io"
	"os"
	"time"

	"github.com/spetix/days2xmasleft/internal/dateutil"
)

type RawOut struct {
	baseOutput
}

func NewRawOut(device *os.File) *RawOut {
	return &RawOut{
		baseOutput: baseOutput{
			Device: device,
		},
	}
}

func (r *RawOut) Print(timeToXmas time.Duration, renderOptions *RenderOptions) {

	data := dateutil.Format(timeToXmas, time.Duration(renderOptions.Unit))

	writer := io.Writer(r.Device)
	//defer io.WriteCloser(r.Device).Close()

	// long and short format are identical

	// writer.Write([]byte(renderOptions.Label))
	writer.Write([]byte(data))
	writer.Write([]byte("\n"))
	// writer.Write([]byte(renderOptions.Label))
	writer.Write([]byte(data))
	writer.Write([]byte("\n"))

	// if renderOptions.ForegroundColor != "" {
	// 	writer.Write([]byte(renderOptions.ForegroundColor))
	// 	writer.Write([]byte("\n"))
	// }
	// if renderOptions.BackgroundColor != "" {
	// 	writer.Write([]byte(renderOptions.BackgroundColor))
	// 	writer.Write([]byte("\n"))
	// }
}
