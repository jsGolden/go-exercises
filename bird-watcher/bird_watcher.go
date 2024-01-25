package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	result := 0
	for i := 0; i < len(birdsPerDay); i++ {
		result += birdsPerDay[i]
	}
	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week
/*
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	BirdsInWeek(birdsPerDay, 2)
	// => 12
*/
func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := (week - 1) * 7
	endIndex := week * 7
	if startIndex < 0 || endIndex > len(birdsPerDay) {
		return 0
	}

	return TotalBirdCount(birdsPerDay[startIndex:endIndex])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
