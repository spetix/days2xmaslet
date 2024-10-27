package blockletapi

import (
	"os"
	"time"

	"github.com/spetix/days2xmasleft/internal/dateutil"
)

type BlockletImpl struct {
	renderOptions *RenderOptions
	out           BlockletOutput
}

func getOut(protocol string) BlockletOutput {

	switch protocol {
	case "json":
		return NewJsonOut(os.Stdout)
	case "raw":
		return NewRawOut(os.Stdout)
	default:
		return NewRawOut(os.Stdout)
	}
}

func New(renderOptions *RenderOptions, proto string) Blocklet {
	return &BlockletImpl{
		renderOptions: renderOptions,
		out:           getOut(proto),
	}
}

func (b *BlockletImpl) Print() {
	b.out.Print(dateutil.HowManyDaysToXmas(time.Now()), b.renderOptions)
}
