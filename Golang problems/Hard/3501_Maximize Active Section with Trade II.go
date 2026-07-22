package main

import "sort"

type tradeRun struct {
	oneStart  int
	oneEnd    int
	leftStart int
	rightEnd  int
	gain      int
}

type maxTree struct {
	size int
	data []int
}

func newMaxTree(values []int) maxTree {
	size := 1
	for size < len(values) {
		size <<= 1
	}

	tree := maxTree{
		size: size,
		data: make([]int, size*2),
	}
	copy(tree.data[size:], values)
	for i := size - 1; i > 0; i-- {
		if tree.data[i<<1] > tree.data[i<<1|1] {
			tree.data[i] = tree.data[i<<1]
		} else {
			tree.data[i] = tree.data[i<<1|1]
		}
	}
	return tree
}

func (tree maxTree) query(left, right int) int {
	best := 0
	for left, right = left+tree.size, right+tree.size; left < right; left, right = left>>1, right>>1 {
		if left&1 == 1 {
			if tree.data[left] > best {
				best = tree.data[left]
			}
			left++
		}
		if right&1 == 1 {
			right--
			if tree.data[right] > best {
				best = tree.data[right]
			}
		}
	}
	return best
}

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	totalActive := 0
	for i := range s {
		if s[i] == '1' {
			totalActive++
		}
	}

	type run struct {
		start int
		end   int
		value byte
	}
	runs := make([]run, 0)
	for start := 0; start < len(s); {
		end := start + 1
		for end < len(s) && s[end] == s[start] {
			end++
		}
		runs = append(runs, run{start: start, end: end - 1, value: s[start]})
		start = end
	}

	candidates := make([]tradeRun, 0)
	for i := 1; i+1 < len(runs); i++ {
		if runs[i].value != '1' || runs[i-1].value != '0' || runs[i+1].value != '0' {
			continue
		}
		candidates = append(candidates, tradeRun{
			oneStart:  runs[i].start,
			oneEnd:    runs[i].end,
			leftStart: runs[i-1].start,
			rightEnd:  runs[i+1].end,
			gain:      runs[i].start - runs[i-1].start + runs[i+1].end - runs[i].end,
		})
	}

	gains := make([]int, len(candidates))
	for i, candidate := range candidates {
		gains[i] = candidate.gain
	}
	tree := newMaxTree(gains)

	answer := make([]int, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		first := sort.Search(len(candidates), func(i int) bool {
			return candidates[i].oneStart > left
		})
		last := sort.Search(len(candidates), func(i int) bool {
			return candidates[i].oneEnd >= right
		})

		bestGain := 0
		if first < last {
			if first+1 == last {
				candidate := candidates[first]
				leftZeroStart := candidate.leftStart
				if left > leftZeroStart {
					leftZeroStart = left
				}
				rightZeroEnd := candidate.rightEnd
				if right < rightZeroEnd {
					rightZeroEnd = right
				}
				bestGain = candidate.oneStart - leftZeroStart + rightZeroEnd - candidate.oneEnd
			} else {
				firstCandidate := candidates[first]
				leftZeroStart := firstCandidate.leftStart
				if left > leftZeroStart {
					leftZeroStart = left
				}
				bestGain = firstCandidate.oneStart - leftZeroStart + firstCandidate.rightEnd - firstCandidate.oneEnd

				lastCandidate := candidates[last-1]
				rightZeroEnd := lastCandidate.rightEnd
				if right < rightZeroEnd {
					rightZeroEnd = right
				}
				lastGain := lastCandidate.oneStart - lastCandidate.leftStart + rightZeroEnd - lastCandidate.oneEnd
				if lastGain > bestGain {
					bestGain = lastGain
				}

				if first+1 < last-1 {
					middleGain := tree.query(first+1, last-1)
					if middleGain > bestGain {
						bestGain = middleGain
					}
				}
			}
		}
		answer[i] = totalActive + bestGain
	}
	return answer
}
