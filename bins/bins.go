package bins

import (
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createAt"`
	Name      string    `json:"name"`
}

var BinList []Bin

func main() {

}

func NewBin(id, name string, private bool, createdAt time.Time) *Bin {
	return &Bin{
		Id:        id,
		Name:      name,
		Private:   private,
		CreatedAt: createdAt,
	}
}
