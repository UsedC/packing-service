package calculator

// CalculatePacks calculates the number of packs needed for a given total and pack sizes. Pack sizes must be sorted in ascending order.
// Returns a nil map if the total cannot be made up by the pack sizes.
func CalculatePacks(total int, packSizes []int) map[int]int {
	if total == 0 {
		return make(map[int]int)
	}
	if !validPackSizes(packSizes) {
		return nil
	}

	maxSize := packSizes[len(packSizes)-1]
	upperBound := total + maxSize

	minNumPacksForSum := make([]int, upperBound+1)
	lastUsedPackForSum := make([]int, upperBound+1)

	for s := 1; s <= upperBound; s++ {
		minNumPacksForSum[s] = upperBound + 1
		lastUsedPackForSum[s] = -1
	}

	minNumPacksForSum[0] = 0

	for sum := 0; sum <= upperBound; sum++ {
		if minNumPacksForSum[sum] == upperBound+1 {
			continue
		}

		currentPacksCount := minNumPacksForSum[sum]

		for _, packSize := range packSizes {
			nextSum := sum + packSize
			if nextSum <= upperBound {
				if currentPacksCount+1 < minNumPacksForSum[nextSum] {
					minNumPacksForSum[nextSum] = currentPacksCount + 1
					lastUsedPackForSum[nextSum] = packSize
				}
			}
		}
	}

	bestSum := -1
	for sum := total; sum <= upperBound; sum++ {
		if minNumPacksForSum[sum] != upperBound+1 {
			bestSum = sum
			break
		}
	}

	result := make(map[int]int)

	for bestSum > 0 {
		packSize := lastUsedPackForSum[bestSum]
		result[packSize]++
		bestSum -= packSize
	}

	return result
}

func validPackSizes(packSizes []int) bool {
	if len(packSizes) == 0 {
		return false
	}

	for i := 1; i < len(packSizes); i++ {
		if packSizes[i] <= packSizes[i-1] || packSizes[i] <= 0 {
			return false
		}
	}

	return true
}
