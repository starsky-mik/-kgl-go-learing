package iteration

func Repeat(word string, count int) string {
	if word == "" || count <= 0 {
		return ""
	}

	if count == 1 {
		return word
	}

	var result string
	for i := 0; i < count; i++ {
		result += word
	}

	return result
}
