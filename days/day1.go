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
