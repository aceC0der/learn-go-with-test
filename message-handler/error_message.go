package message

import "testing"

const red = "\033[31m"
const reset = "\033[0m"

func PrintError(t *testing.T, msg *string) {
	t.Errorf("%s%s\n%s", red, *msg, reset)
}