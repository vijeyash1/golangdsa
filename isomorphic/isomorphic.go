package isomorphic

import "fmt"

// isomorphic means "taa" can be replaced with "ecc".
// all occurances of a char is replaced with another char
// same replaced char cannot be used to replace another diff char

func FindIsomorphic(word, target string) {
	if len(word) != len(target) {
		fmt.Println("Its not isomorphic")
		return
	}
	t := []rune(target)
	find := make(map[rune]rune)
	for i, ch := range word {
		if _, ok := find[ch]; !ok {
			find[ch] = t[i]
			continue
		} else {
			if find[ch] != t[i] {
				fmt.Println("Its not isomorphic")
				return
			}
		}

	}
	fmt.Println("Its isomorphic")
}
