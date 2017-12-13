package algorithms

func ReverseString(word string) string {
	var reversed string
	for _, v := range word {
		reversed = string(v) + reversed
	}
	return reversed
}
