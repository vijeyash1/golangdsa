package anagram

import "fmt"

// anagram is word formed for rearranging exactly the same word using same letters.

func FindAnagram(word, target string) {
	if len(word) != len(target) {
		fmt.Println("its not a anagram")
		return
	}
	t := make(map[rune]int)
	for _, v := range word {
		t[v]++
	}
	for _, w := range target {
		t[w]--
		if t[w] < 0 {
			fmt.Println("its not a anagram")
			return
		}
	}

	fmt.Println("its a anagram")
}
