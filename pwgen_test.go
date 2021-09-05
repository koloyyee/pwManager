package main

import (
	"testing"
)

func TestPasswordGenerator(t *testing.T) {
	t.Run("length over 99", func(t *testing.T) {
		inputs := UserInput{"David", 100, true, true, true}

		got := CheckLen(inputs.Length)
		want := ErrLengthInvalid

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("length less than 8", func(t *testing.T) {
		inputs := UserInput{"David", 7, true, true, true}

		got := CheckLen(inputs.Length)
		want := ErrLengthInvalid
		assertErrMsg(t, got, want)
	})

	t.Run("Same password", func(t *testing.T) {
		inputs := UserInput{"David", 12, true, true, true}
		got := inputs.PasswordGenerator()
		want := inputs.PasswordGenerator()

		if got == want {
			t.Error("The password is not a randomly generated.")
		}
	})
}

func assertErrMsg(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func BenchmarkPasswordGenerator(b *testing.B) {
	inputs := UserInput{"David", 9, true, true, true}
	for i := 0; i < b.N; i++ {
		inputs.PasswordGenerator()
	}
}
