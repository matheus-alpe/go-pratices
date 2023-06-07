package tour

import "fmt"

func init() {
	fmt.Println(Hello("Matheus"))
	PrintRandNumber()
	fmt.Println("secret", SecretValue)
	fmt.Println(Add(2, 5))
	fmt.Println(Swap("Matheus", "Thiago"))
	fmt.Println(Split(20))
	PrintVars()
}
