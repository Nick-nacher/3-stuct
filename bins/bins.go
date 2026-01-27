package bins

import (
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

var BinList []Bin

func main() {

}

func NewBin(id, name string, private bool, createdAt time.Time) *Bin {
	return &Bin{
		id:        id,
		name:      name,
		private:   private,
		createdAt: createdAt,
	}
}
