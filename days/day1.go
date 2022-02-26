package days

func FirstPart(depths []int) int {
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			count++
		}
	}
	return count
}

func SecondPart(depths []int) int {
	count := 0
	for i := 0; i < len(depths)-3; i++ {
		firstWindow := depths[i] + depths[i+1] + depths[i+2]
		nextWindow := depths[i+1] + depths[i+2] + depths[i+3]
		if firstWindow < nextWindow {
			count++
		}
	}
	return count
}
