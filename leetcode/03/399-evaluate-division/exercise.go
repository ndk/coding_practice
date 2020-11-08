package exercise

type relation struct {
	name  string
	value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string][]relation{}
	for i := 0; i < len(equations); i++ {
		graph[equations[i][0]] = append(graph[equations[i][0]], relation{equations[i][1], values[i]})
		graph[equations[i][1]] = append(graph[equations[i][1]], relation{equations[i][0], 1 / values[i]})
	}

	var dfs func(visited map[string]struct{}, source string, target string, value float64) float64
	dfs = func(visited map[string]struct{}, source string, target string, value float64) float64 {
		if _, ok := visited[source]; ok {
			return -1
		}

		if source == target {
			if _, ok := graph[source]; !ok {
				return -1.0
			}
			return value
		}

		visited[source] = struct{}{}

		for _, next := range graph[source] {
			if result := dfs(visited, next.name, target, value*next.value); result != -1.0 {
				return result
			}
		}

		return -1.0
	}

	result := make([]float64, len(queries))
	for i, query := range queries {
		visited := map[string]struct{}{}
		result[i] = dfs(visited, query[0], query[1], 1)
	}
	return result
}
