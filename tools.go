package kalkan

func byteSlice(data []byte) []byte {
	for i, v := range data {
		if v == 0 {
			return data[:i]
		}
	}
	return data
}
