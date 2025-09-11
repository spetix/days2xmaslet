package blockletapi

import (
	"github.com/spetix/bar-out-adapters/pkg/barout"
	"github.com/spetix/bar-out-adapters/pkg/barout/data"
)

type BlockletImpl struct {
	data data.Data
	out  barout.BlockletOutput
}

func New(data data.Data, proto string) Blocklet {
	return &BlockletImpl{
		data: data,
		out:  barout.New(proto),
	}
}

func (b *BlockletImpl) Print() {
	//b.out.Print(dateutil.HowManyDaysToXmas(time.Now(), b.renderOptions.Unit), b.renderOptions)
	b.out.Print(b.data)
}
