package message

import (
	"fmt"
	"testing"
)

const red = "\033[31m"
const reset = "\033[0m"

func PrintError(t *testing.T, got *string, want *string) {
	errorMessage := fmt.Sprintf("got: %s want: %s\n", *got, *want);
	t.Errorf("%s%s\n%s", red, errorMessage, reset)
}

func PrintErrorMessage(t testing.TB, errorMessage *string){
	t.Helper()
	t.Errorf("%s%s\n%s", red, *errorMessage, reset)
}