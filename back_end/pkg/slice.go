package pkg

import "github.com/gocql/gocql"

func RemoveFromSlice(slice []int, element int) []int {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func RemoveSubset(originalSlice []gocql.UUID, subset []gocql.UUID) []gocql.UUID {
	// Create a map to store the elements in the subset
	subsetMap := make(map[gocql.UUID]bool)
	for _, value := range subset {
		subsetMap[value] = true
	}

	// Create a new slice to store the result
	resultSlice := make([]gocql.UUID, 0, len(originalSlice))

	// Iterate over the original slice and append elements that are not in the subset
	for _, value := range originalSlice {
		if !subsetMap[value] {
			resultSlice = append(resultSlice, value)
		}
	}

	return resultSlice
}

// func SortByTime([]PostPreview) {

// }
