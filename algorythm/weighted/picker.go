package weighted

import (
	"fmt"
	rand2 "math/rand"
)

type Picker struct {
}

func (p *Picker) RandomPick(items *map[int]interface{}) (interface{}, error) {
	totalWeight := 0
	for w, _ := range *items {
		totalWeight += w
	}

	rand := rand2.Int() % totalWeight
	c := 0
	for w, item := range *items {
		if w > 0 {
			c += w
			if c > rand {
				return item, nil
			}
		}
	}

	return nil, fmt.Errorf("something went wrong")
}

func (p *Picker) Validate(*map[int]interface{}) error {
	return nil
}
