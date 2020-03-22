package main

import (
	"fmt"
)

func main() {
	states := map[string]string{"CA":"Sacramento", "AZ":"Phoenix"}
	capitals := make(map[string]string)
	for key, value := range states {
		capitals[value] = key
	}
	fmt.Println(states)
	fmt.Println(capitals)
}
/* TERMINAL OUTPUT

[akalaj@lifenode $] go run reverseMap.go 
map[AZ:Phoenix CA:Sacramento]
map[Phoenix:AZ Sacramento:CA]
*/
