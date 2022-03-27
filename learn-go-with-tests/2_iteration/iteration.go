package __iteration

func Repeat(character string) string {
	var repeated string
	const count = 5
	for i := 0; i < count; i++ {
		repeated = repeated + character
	}
	return repeated
}
