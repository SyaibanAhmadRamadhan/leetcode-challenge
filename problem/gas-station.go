package problem

func CanCompleteCircuit(gas []int, cost []int) int {
	start := 0
	totalSurplus := 0
	currentSurplus := 0
	for i := range gas {
		currentSurplus += gas[i] - cost[i]
		totalSurplus += gas[i] - cost[i]
		if currentSurplus < 0 {
			start = i + 1
			currentSurplus = 0
		}
	}
	if totalSurplus < 0 {
		return -1
	}

	return start
}
