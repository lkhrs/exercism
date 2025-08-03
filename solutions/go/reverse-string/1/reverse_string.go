package reverse

func Reverse(input string) (output string) {
	for _, v := range input {
		output = string(v) + output
	}
	return
}
