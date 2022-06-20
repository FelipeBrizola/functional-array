package array

func Contains[T any](s []T, callback func(int, T) bool) bool {

	for i, v := range(s) {

		if callback(i, v) {
			return true
		}
	}

	return false
}

func Find[T any](s []T, callback func(int, T) bool) T {

	for i, v := range(s) {

		if callback(i, v) {
			return v
		}	
	}

	return *new(T)
}