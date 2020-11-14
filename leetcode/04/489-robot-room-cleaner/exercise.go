package exercise

var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func dfs(robot *Robot, visited map[[2]int]struct{}, coords [2]int, direction int) {
	visited[coords] = struct{}{}
	robot.Clean()

	for i := 0; i < len(directions); i++ {
		nextDirection := (direction + i) % len(directions)
		nextCoords := [2]int{coords[0] + directions[nextDirection][0], coords[1] + directions[nextDirection][1]}

		if _, ok := visited[nextCoords]; !ok && robot.Move() {
			dfs(robot, visited, nextCoords, nextDirection)

			robot.TurnRight()
			robot.TurnRight()
			robot.Move()
			robot.TurnRight()
			robot.TurnRight()
		}

		robot.TurnRight()
	}
}

func cleanRoom(robot *Robot) {
	dfs(robot, map[[2]int]struct{}{}, [2]int{0, 0}, 0)
}
