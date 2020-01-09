package rand_pick

import (
	"testing"

	"github.com/garupanojisan/rand-pick/algorythm"
)

func TestRandChoice_Pick(t *testing.T) {
	type fields struct {
		Probabilities *map[int]interface{}
		Algorithm     algorythm.Algorithm
	}
	type args struct {
		num int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      map[interface{}]float64
		wantErr   bool
		tolerance float64
	}{
		{
			name: "Algorithm A",
			fields: fields{
				Probabilities: &map[int]interface{}{
					40: "N",
					30: "R",
					20: "SR",
					10: "UR",
				},
				Algorithm: algorythm.NewWeightedPicker(),
			},
			args: args{
				num: 1000000,
			},
			want: map[interface{}]float64{
				"N":  0.4,
				"R":  0.3,
				"SR": 0.2,
				"UR": 0.1,
			},
			wantErr:   false,
			tolerance: 0.001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RandChoice{
				Probabilities: tt.fields.Probabilities,
				Algorithm:     tt.fields.Algorithm,
			}
			got, err := r.Pick(tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			freq := getFrequencyMap(&got)
			for item, f := range freq {
				w := tt.want[item]
				if (w-tt.tolerance) > f || f > (w+tt.tolerance) {
					t.Errorf("Error: out of tolerance")
					return
				}
			}
		})
	}
}

func getFrequencyMap(items *[]interface{}) map[interface{}]float64 {
	res := map[interface{}]float64{}
	count := 0.0
	for _, item := range *items {
		if _, ok := res[item]; ok {
			res[item]++
		} else {
			res[item] = 0
		}
		count++
	}
	for item, c := range res {
		res[item] = c / count
	}
	return res
}
