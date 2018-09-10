package Algorithm

func UniqueMorseCodeWords (words []string) int {
	encoding := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..",
	"--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	var encodingWords []string
	for _,v1 := range words {
		var s string
		for _,v2 := range []byte(v1) {
			s += encoding[int(v2)-97]
		}
		encodingWords = append(encodingWords, s)
	}
	diff := make(map[string]int)
	for _,k := range encodingWords {
		diff[k]++
	}

	return len(diff)
}
