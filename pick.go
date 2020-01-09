package rand_pick

import "github.com/garupanojisan/rand-pick/algorythm"

type RandChoice struct {
	Probabilities *map[int]interface{}
	Algorithm     algorythm.Algorithm
}

func (r *RandChoice) Pick(num int) ([]interface{}, error) {
	if err := r.Algorithm.Validate(r.Probabilities); err != nil {
		return nil, err
	}
	items := make([]interface{}, num)
	for i := 0; i < num; i++ {
		item, err := r.Algorithm.RandomPick(r.Probabilities)
		if err != nil {
			return nil, err
		}
		items[i] = item
	}
	return items, nil
}
