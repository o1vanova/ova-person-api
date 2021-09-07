package utils

// GetSplitSliceBySize This function converts a slice into a slice of slices (chunks of the same size)
func GetSplitSliceBySize(slice []int, size int) [][]int {
	if size < 1 {
		return nil
	}

	len := len(slice)
	count := len / size
	slices := make([][]int, 0, size)

	for i := 0; i < size; i++ {
		from := i * count
		slices = append(slices, slice[from:from+count])
	}

	if len%size > 0 {
		slices[size-1] = append(slices[size-1], slice[size*count:]...)
	}
	return slices
}

// GetMapWithReversedKey This function converts map ("key-value") into map ("value-key")
func GetMapWithReversedKey(oldMap map[int]string) map[string]int {
	newMap := make(map[string]int, len(oldMap))
	for key, value := range oldMap {
		if newMap[value] != 0 {
			panic("This key was already added")
		}
		newMap[value] = key
	}
	return newMap
}

// GetFilterSliceByExcludeSlice This function returns the filtered slice without excluded elements
func GetFilterSliceByExcludeSlice(slice []int, excludedElements []int) []int {
	if len(excludedElements) < 1 {
		return slice
	}

	set := make(map[int]struct{})
	var newSlice []int

	for _, value := range excludedElements {
		if _, ok := set[value]; !ok {
			set[value] = struct{}{}
		}
	}
	for _, value := range slice {
		if _, ok := set[value]; !ok {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}
