package iteration

func Repeat(characters string, times int) string {
	var result = ""
	for i := 0; i < times; i++ {
		result += characters
	}
	return result
}
