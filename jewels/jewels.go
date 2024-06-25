package jewels

import "fmt"

//we will be given stones that are jewels and stones which represents
//  all stones we have. we need to find how many jewels we have

func FindJewels(jewels string, stones string) {
	count := 0
	for _, j := range jewels {
		for _, s := range stones {
			if j == s {
				count++
			}
		}
	}
	fmt.Printf("total jewels we have is %d", count)
}
