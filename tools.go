package kalkan

// Trims the slice.
func trimSlice(data []byte) []byte {
	for i, v := range data {
		if v == 0 {
			return data[:i]
		}
	}
	return data
}
