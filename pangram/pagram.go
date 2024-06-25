package pangram

import (
	"fmt"
	"strings"
)

// pangram:
// is a sentence that has all the characters of english alphabets showing atleast once

func FindPangram(sentence string) {
	// thequickbrownfoxjumpsoveralazydog
	lower := strings.ToLower(sentence)
	alpha := make(map[rune]int)

	for _, v := range lower {
		alpha[v]++
	}

	if len(alpha) != 26 {
		fmt.Printf("the sentence \"%v\" is not a pangram\n", lower)
	} else {
		fmt.Printf("the sentence \"%v\" is a pangram\n", lower)

	}
}
