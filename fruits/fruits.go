package fruits

import (
	"fmt"
	"strings"
)

func FindFruit() {
	fruits := map[int]string{
		97:  "apple",                // a
		98:  "banana",               // b
		99:  "cherry",               // c
		100: "date",                 // d
		101: "elderberry",           // e
		102: "fig",                  // f
		103: "grape",                // g
		104: "honeydew",             // h
		105: "kiwi",                 // i
		106: "jackfruit",            // j
		107: "kiwano",               // k
		108: "lemon",                // l
		109: "mango",                // m
		110: "nectarine",            // n
		111: "orange",               // o
		112: "papaya",               // p
		113: "quince",               // q
		114: "raspberry",            // r
		115: "strawberry",           // s
		116: "tangerine",            // t
		117: "ugli fruit",           // u
		118: "vanilla",              // v
		119: "watermelon",           // w
		120: "xigua",                // x (another name for watermelon)
		121: "yellow passion fruit", // y
		122: "zucchini",             // z (technically a fruit in the botanical sense)
	}

	for {
		fmt.Println("Enter the fruit name to find")

		var input string
		fmt.Scan(&input)

		input = strings.ToLower(input)
		find := int(input[0])
		if v, ok := fruits[find]; ok {
			if v == input {
				fmt.Printf("the fruit %v was found \n", v)
			} else {
				fmt.Printf("the fruit %v was not found \n", input)
				return
			}
		} else {
			fmt.Printf("the fruit %v was not found \n", input)
			return
		}
	}

}
