package algorythm

import "github.com/garupanojisan/rand-pick/algorythm/weighted"

func NewWeightedPicker() Algorithm {
	return &weighted.Picker{}
}
