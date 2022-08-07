package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	commands := []string{"DetectSquares", "add", "add", "add", "count", "count", "add", "count"}
	points := [][]int{{}, {3, 10}, {11, 2}, {3, 2}, {11, 10}, {14, 8}, {11, 2}, {11, 10}}
	for idx := 0; idx < b.N; idx++ {
		RunDetectSquares(commands, points)
	}
}
func TestRunDetectSquares(t *testing.T) {
	type args struct {
		commands []string
		points   [][]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example1",
			args: args{
				commands: []string{"DetectSquares", "add", "add", "add", "count", "count", "add", "count"},
				points:   [][]int{{}, {3, 10}, {11, 2}, {3, 2}, {11, 10}, {14, 8}, {11, 2}, {11, 10}},
			},
			want: []string{"null", "null", "null", "null", "1", "0", "null", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunDetectSquares(tt.args.commands, tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunDetectSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
