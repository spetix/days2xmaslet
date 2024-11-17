package blockletapi

import (
	"os"
	"time"

	"github.com/spetix/days2xmasleft/internal/dateutil"
	"github.com/spetix/days2xmasleft/internal/popup"
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

func (b *BlockletImpl) Execute() {

	if b.renderOptions.ButtonClicked > 0 {
		b.doPopUP()
		return
	}

	b.out.Print(dateutil.HowManyDaysToXmas(time.Now(), b.renderOptions.Unit), b.renderOptions)
}

func (b *BlockletImpl) doPopUP() {
	p := popup.New()
	p.Execute()
}
