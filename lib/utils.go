package lib

import "fmt"

// PrintFavorites print all items
func PrintFavorites() {
	for _, v := range favorites {
		fmt.Println(v)
	}
}
