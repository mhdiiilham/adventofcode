package main

import "testing"

func TestNumberOfMeasurementIncreased(t *testing.T) {
	type args struct {
		inputs []int
	}
	tests := []struct {
		name               string
		args               args
		wantIncreasedNTime int
	}{
		{
			name: "increased 7 times",
			args: args{
				inputs: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			wantIncreasedNTime: 7,
		},
		{
			name: "input zero",
			args: args{
				inputs: []int{},
			},
			wantIncreasedNTime: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIncreasedNTime := NumberOfMeasurementIncreased(tt.args.inputs); gotIncreasedNTime != tt.wantIncreasedNTime {
				t.Errorf("NumberOfMeasurementIncreased() = %v, want %v", gotIncreasedNTime, tt.wantIncreasedNTime)
			}
		})
	}
}
