package exercise

func findBusiestPeriod(data [][3]int) int {
	busiest := 0
	max := 0
	amount := 0
	for i, snapshot := range data {
		if snapshot[2] == 0 {
			amount -= snapshot[1]
		} else {
			amount += snapshot[1]
		}
		if i < len(data)-1 && data[i+1][0] == snapshot[0] {
			continue
		}
		if amount > max {
			busiest = snapshot[0]
			max = amount
		}
	}
	return busiest
}
