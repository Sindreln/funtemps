package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{input: "sun", want: []string{
			"The Sun is over 4.5 billion years old.",
			"The Sun's light reaches the Earth in eight minutes.",
			"The Sun is made from Hydrogen and Helium.",
		}},
		{input: "luna", want: []string{
			"The Moon always shows Earth the same face.",
			"The Moon's surface is actually dark.",
			"The Moon has quakes too.",
		}},
		{input: "terra", want: []string{
			"Earth has never been perfectly round",
			"Earth's gravity isn't uniform.",
			"In the past, sea levels were very different.",
		}},
		{input: "invalid", want: []string{}},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
