package helpers

func Contains(arr []string, val string) bool {
	for _, element := range arr {
		if element == val {
			return true
		}
	}
	return false
}
