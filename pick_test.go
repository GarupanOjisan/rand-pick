package rand_pick

import (
	"github.com/garupanojisan/rand-pick/algorythm"
	"reflect"
	"testing"
)

func TestNewRandChoice(t *testing.T) {
	tests := []struct {
		name string
		want *RandChoice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRandChoice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandChoice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandChoice_Pick(t *testing.T) {
	type fields struct {
		Probabilities *map[float64]interface{}
		Algorithm     algorythm.Algorithm
	}
	type args struct {
		num int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "Algorithm A",
			fields: fields{
				Probabilities: &map[float64]interface{}{
					0.9: "a",
					0.1: "b",
				},
				Algorithm: algorythm.NewAPicker(),
			},
			args: args{
				num: 1,
			},
			want:    []interface{}{"a", "b"},
			wantErr: false,
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
			if int64(len(got)) != tt.args.num {
				t.Errorf("Pick() len(got) = %d, len(want) %d", len(got), tt.args.num)
				return
			}

			ok := false
			for _, w := range tt.want {
				if reflect.DeepEqual(got, w) {
					ok = true
					break
				}
			}
			if !ok {
				t.Errorf("Pick() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
