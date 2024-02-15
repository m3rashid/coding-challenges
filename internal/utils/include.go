package utils

func Includes[T interface{}](element *T, array []*T) bool {
	for _, arrayEl := range array {
		if element == arrayEl {
			return true
		}
	}

	return false
}
