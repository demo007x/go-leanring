package integers

func Add(x, y int) int {
	return x + y
}

func Repeat(character string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated += character
	}

	return repeated
}
