package blockletapi

import (
	"github.com/spetix/days2xmasleft/internal/data"

	"github.com/spetix/days2xmasleft/internal/out"
)

type BlockletImpl struct {
	data data.Data
	out  out.BlockletOutput
}

func New(data data.Data, proto string) Blocklet {
	return &BlockletImpl{
		data: data,
		out:  out.New(proto),
	}
}

func (b *BlockletImpl) Print() {
	//b.out.Print(dateutil.HowManyDaysToXmas(time.Now(), b.renderOptions.Unit), b.renderOptions)
	b.out.Print(b.data)
}
