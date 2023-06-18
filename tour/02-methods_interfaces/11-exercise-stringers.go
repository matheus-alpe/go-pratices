package methodsinterfaces

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	s := make([]string, len(ip))

	for i, b := range ip {
		s[i] = strconv.Itoa(int(b))
	}

	return strings.Join(s, ".")
}

func ExerciseStringersExample() {
	fmt.Println("\nExercise Stringers:")

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
