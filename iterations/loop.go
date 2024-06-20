package iterations

const character = "a"
const repeatCount = 5

func Loop(character string, repeatCount int) string {
	repeated := ""
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

