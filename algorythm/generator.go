package algorythm

import (
	"github.com/garupanojisan/rand-pick/algorythm/a"
)

type Algorithm interface {
	RandomPick(int64, *map[float64]interface{}) ([]interface{}, error)
	Validate(*map[float64]interface{}) error
}

func NewAPicker() Algorithm {
	return &a.APicker{}
}
