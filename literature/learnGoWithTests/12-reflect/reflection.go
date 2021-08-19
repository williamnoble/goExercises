package reflect

func walk(_ interface{}, fn func(input string)) {
	fn("I still don't believe it")
}
