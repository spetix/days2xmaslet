package blockletapi

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spetix/days2xmasleft/internal/dateutil"
)

type WaybarOut struct {
	baseOutput
}
type outputjsonWaybar struct {
	Text            string `json:"text"`
	Tooltip         string `json:"tooltip"`
	Alt             string `json:"alt"`
	BackgroundColor string `json:"background-color"`
	ForegroundColor string `json:"foreground-color"`
}

func NewWaybarOut(device *os.File) BlockletOutput {
	return &WaybarOut{
		baseOutput: baseOutput{
			Device: device,
		},
	}
}

func (j *WaybarOut) Print(unit time.Duration, rerenderOptions *RenderOptions) {
	newJson := outputjsonWaybar{
		Alt:             rerenderOptions.Label,
		Text:            dateutil.Format(unit, rerenderOptions.Unit),
		Tooltip:         fmt.Sprint(dateutil.Format(unit, rerenderOptions.Unit), " to xmas"),
		BackgroundColor: rerenderOptions.BackgroundColor,
		ForegroundColor: rerenderOptions.ForegroundColor,
	}
	b, err := json.Marshal(newJson)
	if err != nil {
		log.Print("json marshal error", err)
	}
	j.Device.Write(b)
}
