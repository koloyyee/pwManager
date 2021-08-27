package main

import (
	p "github.com/koloyyee/pwgen/generator"
	v "github.com/koloyyee/pwgen/validation"
)

func main() {
	inputs := v.Input()
	length, upper, symbols := v.InputValidator(v.Inputs{inputs.Length, inputs.Uppercase, inputs.Symbols})
	tempPW, pwLength := p.Choice(length, upper, symbols)
	p.PWGenerator(tempPW, pwLength)
}
