package values

import "fmt"

func Value() {
	fmt.Println("go" + "lang") // Strings can be added together with +

	fmt.Println("1+1 =", 1+1)
	fmt.Println("8.0/3.0 =", 8.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
