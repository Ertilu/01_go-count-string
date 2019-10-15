package main

import (
	"testing"
)

func TestCount(t *testing.T) {
	// test is the input is nil
	t.Run("Nil input", func(t *testing.T) {
		input := ""
		consonants, vowels := count(input)
		expectedConsonant := "Huruf mati: 0,"
		expectedVowels := "Huruf hidup: 0"
		if consonants != expectedConsonant && vowels != expectedVowels {
			t.Fatalf("return must be %s %s instead of Huruf mati: %s Huruf hidup: %s", expectedConsonant, expectedVowels, consonants, vowels)
		}
	})
}
