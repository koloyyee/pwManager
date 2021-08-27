package generator

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	LowerChar = "abcdefghjklmnopqrstuvwxyz"
	UpperChar = "ABCDEFGHJKLMNOPQRSTUVWXYZ"
	Num       = "1234567890"
	Symbols   = "`-=[]\\;',./~!@#$%^&*()_+{}|:\"<>?"
)

// PWGenerator generates the final outcome of the random password.
func PWGenerator(tempPW []string, length int) string {
	// randPW store random ordered character for password.
	randPW := []string{}
	for i := 0; i < length; i++ {
		randInt := rand.Intn(len(tempPW))
		randPW = append(randPW, tempPW[randInt])
	}
	yourPW := strings.Join(randPW, "")
	fmt.Println(yourPW)
	return yourPW
}

// Choice check the input by client is y or n which leads password features uppercase and/or symbols.
func Choice(length int, upper, symbols string) ([]string, int) {
	var tempPW []string
	switch {
	case upper == "n" && symbols == "n":
		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(LowerChar[rand.Intn(len(LowerChar))]))
		}
		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(Num[rand.Intn(len(Num))]))
		}
	case upper == "y" && symbols == "n":
		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(LowerChar[rand.Intn(len(LowerChar))]))
		}
		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(UpperChar[rand.Intn(len(UpperChar))]))
		}
		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(Num[rand.Intn(len(Num))]))
		}
	case upper == "n" && symbols == "n":
		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(LowerChar[rand.Intn(len(LowerChar))]))
		}
		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(Num[rand.Intn(len(Num))]))
		}

	default:
		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(UpperChar[rand.Intn(len(UpperChar))]))
		}
		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(LowerChar[rand.Intn(len(LowerChar))]))
		}

		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(Num[rand.Intn(len(Num))]))
		}

		for i := 0; i < length; i++ {
			tempPW = append(tempPW, string(Symbols[rand.Intn(len(Symbols))]))
		}
	}

	return tempPW, length
}
