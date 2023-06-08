package basics

import (
	"fmt"
)

func init() {
	fmt.Println(Hello("Matheus"))
	PrintRandNumber()
	fmt.Println("secret", SecretValue)
	fmt.Println(Add(2, 5))
	fmt.Println(Swap("Matheus", "Thiago"))
	fmt.Println(Split(20))
	PrintVars()
	PrintBasicTypes()
	PrintZeroValues()
	PrintTypeConversion()
	PrintTypeInference()
	fmt.Println(NeedInt(Small))
	fmt.Println(NeedFloat(Small))
	fmt.Println(NeedFloat(Big))
	ForExample()
	ForContinuedExample()
	ForWhileExample()
	IfExample()
}
