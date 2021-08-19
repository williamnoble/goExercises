package iteration

func Repeat(character string, repeat int) (repeated string) {

	// Option 1
	for i := 0; i < repeat; i++ {
		repeated += character
	}
	return repeated

	// Option 2
	//return strings.Repeat(character, 5)
}
