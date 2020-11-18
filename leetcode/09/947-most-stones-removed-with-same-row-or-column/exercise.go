package exercise

func dfs(graph map[int][]int, visited []bool, node int) {
	visited[node] = true
	for _, next := range graph[node] {
		if !visited[next] {
			dfs(graph, visited, next)
		}
	}
}

func removeStones(stones [][]int) int {
	graph := map[int][]int{}
	{
		x, y := map[int][]int{}, map[int][]int{}
		for i, stone := range stones {
			nextStones := graph[i]
			nextStones = append(nextStones, x[stone[0]]...)
			nextStones = append(nextStones, y[stone[1]]...)
			graph[i] = nextStones

			for _, prev := range x[stone[0]] {
				graph[prev] = append(graph[prev], i)
			}
			for _, prev := range y[stone[1]] {
				graph[prev] = append(graph[prev], i)
			}
			x[stone[0]] = append(x[stone[0]], i)
			y[stone[1]] = append(y[stone[1]], i)
		}
	}

	componenets := 0
	{
		visited := make([]bool, len(stones))
		for i := range stones {
			if !visited[i] {
				componenets++
				dfs(graph, visited, i)
			}
		}
	}

	return len(stones) - componenets
}

///////////////////////////////////////////////////////////////////////////////
// The fastest solution on Leetcode at the moment

func removeStones_leet(stones [][]int) int {
	l := len(stones)

	if l == 0 || l == 1 {
		return 0
	}

	parent := make([]int, 20000)
	for i := range parent {
		parent[i] = -1
	}

	for _, s := range stones {
		parent[s[0]] = s[0]
		parent[s[1]+10000] = s[1] + 10000
	}

	for _, s := range stones {
		union(s[0], s[1]+10000, parent)
	}

	numOfUnion := 0
	for i, v := range parent {
		if v == i {
			numOfUnion++
		}
	}

	return l - numOfUnion
}

func removeStones_leet2(stones [][]int) int {

	l := len(stones)
	if l == 0 || l == 1 {
		return 0
	}

	parent := make([]int, l)
	for i := range parent {
		parent[i] = i
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if sameRowOrCol(i, j, stones) {
				union(i, j, parent)
			}
		}
	}

	numOfUnion := 0
	for i, v := range parent {
		if i == v {
			numOfUnion++
		}
	}
	return l - numOfUnion
}

func sameRowOrCol(i, j int, stones [][]int) bool {
	s1 := stones[i]
	s2 := stones[j]

	return s1[0] == s2[0] || s1[1] == s2[1]
}

func find(x int, parent []int) int {
	if x != parent[x] {
		parent[x] = find(parent[x], parent)
	}
	return parent[x]
}

func union(x, y int, parent []int) {
	a := find(x, parent)
	b := find(y, parent)
	parent[a] = b
}
