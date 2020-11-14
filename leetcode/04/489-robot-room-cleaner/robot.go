package exercise

var robotDirections = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Robot struct {
	row, col  int
	direction int
	room      [][]int
	state     [][]int
	remains   int
}

// Returns true if the cell in front is open and robot moves into the cell.
// Returns false if the cell in front is blocked and robot stays in the current cell.
func (robot *Robot) Move() bool {
	newRow, newCol := robot.row+robotDirections[robot.direction][0], robot.col+robotDirections[robot.direction][1]
	if !(0 <= newRow && newRow < len(robot.room) && 0 <= newCol && newCol < len(robot.room[newRow])) {
		return false
	}
	robot.row, robot.col = newRow, newCol
	return true
}

// // Robot will stay in the same cell after calling TurnLeft/TurnRight.
// // Each turn will be 90 degrees.
// func (robot *Robot) TurnLeft() {
// 	robot.direction--
// 	if robot.direction < 0 {
// 		robot.direction = len(robotDirections) - 1
// 	}
// }

func (robot *Robot) TurnRight() {
	robot.direction++
	if robot.direction >= len(robotDirections) {
		robot.direction = 0
	}
}

// Clean the current cell.
func (robot *Robot) Clean() {
	if robot.state[robot.row][robot.col] != 0 {
		robot.state[robot.row][robot.col] = 0
		robot.remains--
	}
}
