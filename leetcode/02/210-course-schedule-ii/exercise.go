package exercise

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := map[int][]int{}
	for _, p := range prerequisites {
		graph[p[0]] = append(graph[p[0]], p[1])
	}

	order := make([]int, numCourses)
	orderPos := 0

	visited := make([]bool, numCourses)
	usage := make([]bool, numCourses)

	var dfs func(i int, orderPos *int)
	dfs = func(i int, orderPos *int) {
		visited[i] = true

		next := graph[i]
		if len(next) == 0 {
			usage[i] = true
			order[*orderPos] = i
			*orderPos++
			return
		}

		for _, j := range graph[i] {
			if visited[j] {
				if !usage[j] {
					return
				}
				continue
			}
			dfs(j, orderPos)
			if !usage[j] {
				return
			}
		}

		usage[i] = true
		order[*orderPos] = i
		*orderPos++
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] {
			if !usage[i] {
				return nil
			}
			continue
		}
		dfs(i, &orderPos)
		if !usage[i] {
			return nil
		}
	}

	return order
}

func findOrderBFS(numCourses int, prerequisites [][]int) []int {
	usage := make([]int, numCourses)
	courses := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		first := prerequisite[0]
		second := prerequisite[1]
		usage[first]++
		courses[second] = append(courses[second], first)
	}

	queue := make([]int, 0, numCourses)
	for course, num := range usage {
		if num == 0 {
			queue = append(queue, course)
		}
	}

	order := make([]int, 0, numCourses)
	for len(queue) != 0 {
		course := queue[0]
		queue = queue[1:]
		order = append(order, course)
		for _, next := range courses[course] {
			usage[next]--
			if usage[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	if len(order) == numCourses {
		return order
	}
	return nil
}
