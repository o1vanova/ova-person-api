package utils

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
