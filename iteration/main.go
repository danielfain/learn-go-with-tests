package iteration

func Repeat(char string, numRepeats int) (repeatedChar string) {
	for i := 0; i < numRepeats; i++ {
		repeatedChar += char
	}
	return
}
