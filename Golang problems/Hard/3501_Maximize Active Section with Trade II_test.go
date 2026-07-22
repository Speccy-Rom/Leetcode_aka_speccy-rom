package main

import (
	"reflect"
	"testing"
)

func TestMaxActiveSectionsAfterTradeExamples(t *testing.T) {
	tests := []struct {
		s       string
		queries [][]int
		want    []int
	}{
		{"01", [][]int{{0, 1}}, []int{1}},
		{"0100", [][]int{{0, 3}, {0, 2}, {1, 3}, {2, 3}}, []int{4, 3, 1, 1}},
		{"1000100", [][]int{{1, 5}, {0, 6}, {0, 4}}, []int{6, 7, 2}},
		{"01010", [][]int{{0, 3}, {1, 4}, {1, 3}}, []int{4, 4, 2}},
	}

	for _, tt := range tests {
		if got := maxActiveSectionsAfterTrade(tt.s, tt.queries); !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("maxActiveSectionsAfterTrade(%q, %v) = %v, want %v", tt.s, tt.queries, got, tt.want)
		}
	}
}

func TestMaxActiveSectionsAfterTradeAgainstBruteForce(t *testing.T) {
	for length := 1; length <= 7; length++ {
		for mask := 0; mask < 1<<length; mask++ {
			bits := make([]byte, length)
			for i := range bits {
				bits[i] = '0' + byte(mask>>i&1)
			}
			s := string(bits)

			queries := make([][]int, 0, length*(length+1)/2)
			for left := 0; left < length; left++ {
				for right := left; right < length; right++ {
					queries = append(queries, []int{left, right})
				}
			}

			got := maxActiveSectionsAfterTrade(s, queries)
			for i, query := range queries {
				want := bruteForceAnswer(s, query[0], query[1])
				if got[i] != want {
					t.Fatalf("s=%q, query=%v: got %d, want %d", s, query, got[i], want)
				}
			}
		}
	}
}

func bruteForceAnswer(s string, left, right int) int {
	totalActive := 0
	for i := range s {
		if s[i] == '1' {
			totalActive++
		}
	}

	augmented := append([]byte{'1'}, s[left:right+1]...)
	augmented = append(augmented, '1')
	bestGain := 0

	for start := 1; start+1 < len(augmented); {
		if augmented[start] != '1' {
			start++
			continue
		}
		end := start
		for end+1 < len(augmented)-1 && augmented[end+1] == '1' {
			end++
		}
		if augmented[start-1] == '0' && augmented[end+1] == '0' {
			afterTrade := append([]byte(nil), augmented...)
			for i := start; i <= end; i++ {
				afterTrade[i] = '0'
			}
			for zeroStart := 1; zeroStart+1 < len(afterTrade); {
				if afterTrade[zeroStart] != '0' {
					zeroStart++
					continue
				}
				zeroEnd := zeroStart
				for zeroEnd+1 < len(afterTrade) && afterTrade[zeroEnd+1] == '0' {
					zeroEnd++
				}
				if afterTrade[zeroStart-1] == '1' && afterTrade[zeroEnd+1] == '1' {
					gain := zeroEnd - zeroStart + 1 - (end - start + 1)
					if gain > bestGain {
						bestGain = gain
					}
				}
				zeroStart = zeroEnd + 1
			}
		}
		start = end + 1
	}
	return totalActive + bestGain
}
