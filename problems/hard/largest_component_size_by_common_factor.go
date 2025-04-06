package hard

// UnionFind represents a disjoint set data structure
type UnionFind struct {
	parent []int
	rank   []int
	size   []int
}

// NewUnionFind creates a new UnionFind structure
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
	}
	for i := range uf.parent {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

// Find returns the root of the set containing x
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y
func (uf *UnionFind) Union(x, y int) {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return
	}
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px
	uf.size[px] += uf.size[py]
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
}

// GetSize returns the size of the set containing x
func (uf *UnionFind) GetSize(x int) int {
	return uf.size[uf.Find(x)]
}

// getPrimeFactors returns all prime factors of a number
func getPrimeFactors(n int) []int {
	factors := make([]int, 0)
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

// largestComponentSize finds the size of the largest connected component
// in the graph where nodes are connected if they share a common factor > 1
func largestComponentSize(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Create a map to store the first occurrence of each prime factor
	factorToIndex := make(map[int]int)
	uf := NewUnionFind(len(nums))

	// For each number, connect it to all its prime factors
	for i, num := range nums {
		factors := getPrimeFactors(num)
		for _, factor := range factors {
			if prevIndex, exists := factorToIndex[factor]; exists {
				uf.Union(i, prevIndex)
			} else {
				factorToIndex[factor] = i
			}
		}
	}

	// Find the largest component size
	maxSize := 0
	for i := range nums {
		size := uf.GetSize(i)
		if size > maxSize {
			maxSize = size
		}
	}

	return maxSize
}
