package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

type UserInput struct {
	Username  string
	Length    int64
	Uppercase bool
	Num       bool
	Symbols   bool
}

const (
	LowerChar = "abcdefghjklmnopqrstuvwxyz"
	UpperChar = "ABCDEFGHJKLMNOPQRSTUVWXYZ"
	Num       = "1234567890"
	Symbols   = "`-=[]\\;',./~!@#$%^&*()_+{}|:\"<>?"
)

// ErrLengthInvalid error message for invalid length input.
var ErrLengthInvalid = errors.New("Length is must between 8-99")

func main() {
	// input from cmd line.

	username := flag.String("name", "", "Your username of the password.")
	length := flag.Int64("len", 0, "The Length of the password between 8-99.")
	uppercase := flag.Bool("upper", true, "Include uppercase in the password?")
	num := flag.Bool("num", true, "Include numbers in the password?")
	symbols := flag.Bool("sym", true, "Include symbols in the password?")

	flag.Parse()

	err := CheckLen(*length)
	if err != nil {
		fmt.Printf("Invalid Input: %v", err)
		os.Exit(1)
	}

	inputs := UserInput{*username, *length, *uppercase, *num, *symbols}
	inputs.PasswordGenerator()
}

// CheckLen make sure the length input is between 8-99.
func CheckLen(len int64) error {
	if len < 8 || len > 99 {
		return ErrLengthInvalid
	}
	return nil
}

// PasswordGenerator will take the inputs from user with the format of UserInput Struct
// and then output the Password requirements format struct
func (u *UserInput) PasswordGenerator() string {
	result := []string{}
	var final string
	switch {
	case u.Uppercase == true && u.Num == true && u.Symbols == true:
		result = u.RandEl(result, UpperChar)
		result = u.RandEl(result, Num)
		result = u.RandEl(result, Symbols)
		result = u.RandEl(result, LowerChar)
	case u.Uppercase == false && u.Num == true && u.Symbols == true:
		result = u.RandEl(result, Num)
		result = u.RandEl(result, Symbols)
		result = u.RandEl(result, LowerChar)
	case u.Uppercase == false && u.Num == false && u.Symbols == true:
		result = u.RandEl(result, Symbols)
		result = u.RandEl(result, LowerChar)
	case u.Uppercase == true && u.Num == true && u.Symbols == false:
		result = u.RandEl(result, UpperChar)
		result = u.RandEl(result, Num)
		result = u.RandEl(result, LowerChar)
	case u.Uppercase == true && u.Num == false && u.Symbols == true:
		result = u.RandEl(result, UpperChar)
		result = u.RandEl(result, Symbols)
		result = u.RandEl(result, LowerChar)
	case u.Uppercase == false && u.Num == true && u.Symbols == true:
		result = u.RandEl(result, Num)
		result = u.RandEl(result, Symbols)
		result = u.RandEl(result, LowerChar)
	case u.Uppercase == true && u.Num == false && u.Symbols == false:
		result = u.RandEl(result, UpperChar)
		result = u.RandEl(result, LowerChar)
	default:
		result = u.RandEl(result, LowerChar)

	}

	if len(result) > 0 {
		final = u.ShufflePW(result)
	}

	fmt.Println(final)
	return final
}

// RandEl Generate elements according to the user's input.
func (u UserInput) RandEl(p []string, element string) []string {
	if u.Length < 8 {
		fmt.Println("Length must be over 8.")
		os.Exit(1)
	}
	for i := 0; i < int(u.Length); i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(element))))
		if err != nil {
			panic(err)
		}

		pI := byte(n.Int64())
		p = append(p, string(element[pI]))
	}
	return p
}

// ShufflePW received a slice of string and reshuffle it to generate a randon password.
func (u UserInput) ShufflePW(p []string) string {
	shuffled := []string{}
	for i := 0; i < int(u.Length); i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(p))))
		if err != nil {
			panic(err)
		}
		pI := byte(n.Int64())
		shuffled = append(shuffled, string(p[pI]))
	}
	final := strings.Join(shuffled, "")
	return final
}
