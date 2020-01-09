package a

type APicker struct {
}

func (a *APicker) RandomPick(num int64, probabilities *map[float64]interface{}) ([]interface{}, error) {
	return nil, nil
}

func (a *APicker) Validate(probabilities *map[float64]interface{}) error {
	return nil
}
