package exercise

func calcDroneMinEnergy(data [][3]int) int {
	if len(data) == 0 {
		return 0
	}

	max := 0
	for i := 1; i < len(data);i++ {
		if data[i][2] > data[max][2] {
			max = i
		}
	}
	return data[max][2]-data[0][2]
}
