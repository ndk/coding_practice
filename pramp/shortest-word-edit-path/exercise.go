package exercise

func shortestWordEditPath(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)

	end := -1

	mutants := map[string][]int{}
	for i, word := range wordList {
		if word == endWord {
			end = i
		}
		for p := range word {
			mutant := word[:p] + "*" + word[p+1:]
			mutants[mutant] = append(mutants[mutant], i)
		}
	}
	graph := map[int]map[int]struct{}{}
	for _, group := range mutants {
		for i := 0; i < len(group); i++ {
			for j := 0; j < len(group); j++ {
				if i == j {
					continue
				}
				source, target := group[i], group[j]
				targets, ok := graph[source]
				if !ok {
					targets = map[int]struct{}{}
				}
				targets[target] = struct{}{}
				graph[source] = targets
			}
		}
	}

	queue := []int{len(wordList) - 1}
	visited := map[int]struct{}{
		queue[0]: {},
	}
	steps := -1
	for len(queue) > 0 {
		steps++
		for size := len(queue); size > 0; size-- {
			current := queue[0]
			if current == end {
				return steps
			}
			queue = queue[1:]
			for next := range graph[current] {
				if _, ok := visited[next]; ok {
					continue
				}
				visited[next] = struct{}{}
				queue = append(queue, next)
			}
		}
	}

	return -1
}
