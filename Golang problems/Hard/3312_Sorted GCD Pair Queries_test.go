package main

import (
	"reflect"
	"testing"
)

func TestGCDValues(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries []int64
		want    []int
	}{
		{
			name:    "mixed gcd values",
			nums:    []int{2, 3, 4},
			queries: []int64{0, 1, 2},
			want:    []int{1, 1, 2},
		},
		{
			name:    "duplicate values",
			nums:    []int{4, 4, 2, 1},
			queries: []int64{0, 1, 5},
			want:    []int{1, 1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdValues(tt.nums, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("gcdValues(%v, %v) = %v, want %v", tt.nums, tt.queries, got, tt.want)
			}
		})
	}
}
