package strutil_test

import (
	"testing"

	"github.com/itispx/cpfcnpj/strutil"
)

func TestEmptyString(t *testing.T) {
	input := ""
	expected := ""

	result := strutil.ExtractNumbers(input)
	if result != expected {
		t.Errorf("expected: %s | received: %s", expected, result)
	}
}

func TestOnlyNumbers(t *testing.T) {
	input := "1234567890"
	expected := "1234567890"

	result := strutil.ExtractNumbers(input)
	if result != expected {
		t.Errorf("expected: %s | received: %s", expected, result)
	}
}

func TestOnlyLetters(t *testing.T) {
	input := "abcdef"
	expected := ""

	result := strutil.ExtractNumbers(input)
	if result != expected {
		t.Errorf("expected: %s | received: %s", expected, result)
	}
}

func TestLettersAndNumbers(t *testing.T) {
	input := "abc123def456"
	expected := "123456"

	result := strutil.ExtractNumbers(input)
	if result != expected {
		t.Errorf("expected: %s | received: %s", expected, result)
	}
}

func TestSpecialCharactersAndNumbers(t *testing.T) {
	input := "!@#123$%^456&*()"
	expected := "123456"

	result := strutil.ExtractNumbers(input)
	if result != expected {
		t.Errorf("expected: %s | received: %s", expected, result)
	}
}

func TestSpacesAndNumbers(t *testing.T) {
	input := "1 2 3 4 5 6"
	expected := "123456"

	result := strutil.ExtractNumbers(input)
	if result != expected {
		t.Errorf("expected: %s | received: %s", expected, result)
	}
}

func TestOnlySpecialCharacters(t *testing.T) {
	input := "!@#$%^&*()"
	expected := ""

	result := strutil.ExtractNumbers(input)
	if result != expected {
		t.Errorf("expected: %s | received: %s", expected, result)
	}
}

func TestMixedContent(t *testing.T) {
	input := "abc123!@#def456$%^ghi789"
	expected := "123456789"

	result := strutil.ExtractNumbers(input)
	if result != expected {
		t.Errorf("expected: %s | received: %s", expected, result)
	}
}
