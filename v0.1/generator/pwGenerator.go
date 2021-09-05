package generator

import (
	"fmt"
	"math/rand"
	"strings"
)

type Password interface {
	PWGenerator() string
	Choice() ([]string, int)
}

type ClientInputs struct {
	Length             int
	Uppercase, Symbols string
}

type PasswordProcess struct {
	TempPW []string
	Length int
}

const (
	LowerChar = "abcdefghjklmnopqrstuvwxyz"
	UpperChar = "ABCDEFGHJKLMNOPQRSTUVWXYZ"
	Num       = "1234567890"
	Symbols   = "`-=[]\\;',./~!@#$%^&*()_+{}|:\"<>?"
)

// PWGenerator generates the final outcome of the random password.
func (p PasswordProcess) PWGenerator() string {
	// randPW store random ordered character for password.
	randPW := []string{}
	for i := 0; i < p.Length; i++ {
		randInt := rand.Intn(len(p.TempPW))
		randPW = append(randPW, p.TempPW[randInt])
	}
	yourPW := strings.Join(randPW, "")
	fmt.Println(yourPW)
	return yourPW
}

// Choice check the input by client is y or n which leads password features uppercase and/or symbols.
func (c ClientInputs) Choice() ([]string, int) {
	var tempPW []string
	switch {
	case c.Uppercase == "n" && c.Symbols == "n":
		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(LowerChar[rand.Intn(len(LowerChar))]))
		}
		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(Num[rand.Intn(len(Num))]))
		}
	case c.Uppercase == "y" && c.Symbols == "n":
		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(LowerChar[rand.Intn(len(LowerChar))]))
		}
		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(UpperChar[rand.Intn(len(UpperChar))]))
		}
		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(Num[rand.Intn(len(Num))]))
		}
	case c.Uppercase == "n" && c.Symbols == "n":
		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(LowerChar[rand.Intn(len(LowerChar))]))
		}
		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(Num[rand.Intn(len(Num))]))
		}

	default:
		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(UpperChar[rand.Intn(len(UpperChar))]))
		}
		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(LowerChar[rand.Intn(len(LowerChar))]))
		}

		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(Num[rand.Intn(len(Num))]))
		}

		for i := 0; i < c.Length; i++ {
			tempPW = append(tempPW, string(Symbols[rand.Intn(len(Symbols))]))
		}
	}

	return tempPW, c.Length
}
