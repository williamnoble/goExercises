package reflect

func walk(x interface{}, fn func(input string)) {
	fn("I still don't believe it")
}

// complete later
