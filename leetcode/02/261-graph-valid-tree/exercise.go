package exercise

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		for ; parent[a] != a; a = parent[a] {
		}
		for ; parent[b] != b; b = parent[b] {
		}
		if a == b {
			return false
		}
		parent[a] = b
	}

	return true
}

///////////////////////////////////////////////////////////////////////////////
// Approach 1: Graph Theory

// Iterative Depth-First Search
func validTree_GT_IDFS(n int, edges [][]int) bool {
	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	parent := map[int]int{
		0: -1,
	}

	stack := []int{0}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, neighbour := range graph[node] {
			if p, ok := parent[node]; ok && p == neighbour {
				continue
			}
			if _, ok := parent[neighbour]; ok {
				return false
			}
			stack = append(stack, neighbour)
			parent[neighbour] = node
		}
	}

	return len(parent) == n
}

// Recursive Depth-First Search
func validTree_GT_RDFS(n int, edges [][]int) bool {
	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	seen := map[int]struct{}{}
	var dfs func(graph map[int][]int, seen map[int]struct{}, node int, parent int) bool
	dfs = func(graph map[int][]int, seen map[int]struct{}, node int, parent int) bool {
		if _, ok := seen[node]; ok {
			return false
		}
		seen[node] = struct{}{}

		if neighbours, ok := graph[node]; ok {
			for _, neighbour := range neighbours {
				if parent != neighbour {
					if !dfs(graph, seen, neighbour, node) {
						return false
					}
				}
			}
		}
		return true
	}
	return dfs(graph, seen, 0, -1) && len(seen) == n
}

// Iterative Breadth-First Search.
func validTree_GT_BFS(n int, edges [][]int) bool {
	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	parent := map[int]int{
		0: -1,
	}

	queue := []int{0}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbour := range graph[node] {
			if p, ok := parent[node]; ok && p == neighbour {
				continue
			}
			if _, ok := parent[neighbour]; ok {
				return false
			}
			queue = append(queue, neighbour)
			parent[neighbour] = node
		}
	}

	return len(parent) == n
}

///////////////////////////////////////////////////////////////////////////////
// Approach 2: Advanced Graph Theory

// Iterative Depth-First Search
func validTree_AGT_IDFS(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	seen := map[int]struct{}{0: {}}
	stack := []int{0}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, neighbour := range graph[node] {
			if _, ok := seen[neighbour]; ok {
				continue
			}
			seen[neighbour] = struct{}{}
			stack = append(stack, neighbour)
		}
	}

	return len(seen) == n
}

// Recursive Depth-First Search
func validTree_AGT_RDFS(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	seen := map[int]struct{}{}
	var dfs func(graph map[int][]int, seen map[int]struct{}, node int)
	dfs = func(graph map[int][]int, seen map[int]struct{}, node int) {
		if _, ok := seen[node]; ok {
			return
		}
		seen[node] = struct{}{}

		if neighbours, ok := graph[node]; ok {
			for _, neighbour := range neighbours {
				dfs(graph, seen, neighbour)
			}
		}
	}
	dfs(graph, seen, 0)
	return len(seen) == n
}

// Iterative Breadth-First Search
func validTree_AGT_BFS(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	queue := []int{0}
	seen := map[int]struct{}{0: {}}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbour := range graph[node] {
			if _, ok := seen[neighbour]; ok {
				continue
			}
			seen[neighbour] = struct{}{}
			queue = append(queue, neighbour)
		}
	}

	return len(seen) == n
}

///////////////////////////////////////////////////////////////////////////////
// Approach 3: Advanced Graph Theory

// Union Find
func validTree_UF(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	find := func(a int) int {
		for ; parent[a] != a; a = parent[a] {
		}
		return a
	}

	union := func(a int, b int) bool {
		rootA, rootB := find(a), find(b)
		if rootA == rootB {
			return false
		}
		parent[rootA] = rootB
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return false
		}
	}

	return true
}

// Union Find + Path Compression
func validTree_UF_PC(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	parent, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parent[i], size[i] = i, 1
	}

	find := func(a int) int {
		root := a
		for ; parent[root] != root; root = parent[root] {
		}
		for a != root {
			a, parent[a] = parent[a], root
		}
		return root
	}

	union := func(a int, b int) bool {
		rootA, rootB := find(a), find(b)
		if rootA == rootB {
			return false
		}
		if size[rootA] < size[rootB] {
			parent[rootA] = rootB
			size[rootB] += size[rootA]
		} else {
			parent[rootB] = rootA
			size[rootA] += size[rootB]
		}
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return false
		}
	}

	return true
}
