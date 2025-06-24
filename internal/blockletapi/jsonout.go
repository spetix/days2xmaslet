package blockletapi

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/spetix/days2xmasleft/internal/dateutil"
)

type JsonOut struct {
	baseOutput
}
type outputjson struct {
	Short           string `json:"short"`
	Long            string `json:"long"`
	Label           string `json:"label"`
	BackgroundColor string `json:"background-color"`
	ForegroundColor string `json:"foreground-color"`
}

func NewJsonOut(device *os.File) BlockletOutput {
	return &JsonOut{
		baseOutput: baseOutput{
			Device: device,
		},
	}
}

func (j *JsonOut) Print(unit time.Duration, rerenderOptions *RenderOptions) {
	newJson := outputjson{
		Label:           rerenderOptions.Label,
		Short:           dateutil.Format(unit, rerenderOptions.Unit),
		Long:            dateutil.Format(unit, rerenderOptions.Unit),
		BackgroundColor: rerenderOptions.BackgroundColor,
		ForegroundColor: rerenderOptions.ForegroundColor,
	}
	b, err := json.Marshal(newJson)
	if err != nil {
		log.Print("json marshal error", err)
	}
	j.Device.Write(b)
}
