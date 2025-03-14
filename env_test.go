package utils_test

import (
	"log"
	"os"

	"testing"

	utils "github.com/alanwade2001/go-sepa-utils"
)

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestEnvIsSet(t *testing.T) {

	key := "KEY1"
	notSet := "NOT_SET"
	set := "SET"

	err := os.Setenv(key, set)
	if err != nil {
		log.Fatalln("failed to set env variable")
	}

	expected := set
	actual := utils.Getenv(key, notSet)

	if actual != set {
		t.Fatalf(`Getenv("") = %q, want "%q", error`, actual, expected)
	}
}

func TestEnvIsNotSet(t *testing.T) {

	key := "KEY1"
	notSet := "NOT_SET"

	expected := notSet
	actual := utils.Getenv(key, notSet)

	if actual != notSet {
		t.Fatalf(`Getenv("") = %q, want "%q", error`, actual, expected)
	}
}
