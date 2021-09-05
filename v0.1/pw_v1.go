package v01

import (
	p "github.com/koloyyee/pwgen/v0.1/generator"
	v "github.com/koloyyee/pwgen/v0.1/validation"
)

func main() {
	rawInputs := v.Input()
	inputs := v.Inputs{Length: rawInputs.Length, Uppercase: rawInputs.Uppercase, Symbols: rawInputs.Symbols}
	length, upper, symbols := inputs.InputValidator()

	clientInputs := p.ClientInputs{Length: length, Uppercase: upper, Symbols: symbols}
	tempPW, pwLength := clientInputs.Choice()

	pw := p.PasswordProcess{TempPW: tempPW, Length: pwLength}
	pw.PWGenerator()
}
