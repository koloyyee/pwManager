package validation

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Validation interface {
	InputValidator() (int, string, string)
	Inputs() Inputs
}

type Inputs struct {
	Length, Uppercase, Symbols string
}

// Input receive inputs from client.
func Input() Inputs {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Length for your password? ")
	l, err := reader.ReadString('\n')
	fmt.Println("Uppercase include? y/n")
	u, err := reader.ReadString('\n')
	fmt.Println("Symbols included? y/n")
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("%e\n", err)
	}
	return Inputs{l, u, s}
}

// InputValidator check the values are correct.
func (i Inputs) InputValidator() (int, string, string) {
	correctInput := false
	var length int
	var upper, symbols string
	for correctInput == false {

		cleanL := strings.Fields(i.Length)
		cleanU := strings.Fields(i.Uppercase)[0]
		cleanS := strings.Fields(i.Symbols)[0]
		lInt, err := strconv.Atoi(cleanL[0])

		if err != nil || cleanU != "n" && cleanU != "y" || cleanS != "n" && cleanS != "y" {
			fmt.Println("Invalid Input, Length must be numeric, Uppercase and Symbols must be y/n.")
			i = Input()
			continue
		}
		length = lInt
		upper = cleanU
		symbols = cleanS
		correctInput = true
	}

	return length, upper, symbols
}
