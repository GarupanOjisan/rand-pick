package rand_pick

import "github.com/garupanojisan/rand-pick/algorythm"

type RandChoice struct {
	Probabilities *map[float64]interface{}
	Algorithm     algorythm.Algorithm
}

func NewRandChoice() *RandChoice {
	return &RandChoice{}
}

func (r *RandChoice) Pick(num int64) ([]interface{}, error) {
	if err := r.Algorithm.Validate(r.Probabilities); err != nil {
		return nil, err
	}
	return r.Algorithm.RandomPick(num, r.Probabilities)
}
