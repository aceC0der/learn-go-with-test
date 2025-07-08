package basic

import (
	"fmt"
	"learn-go-with-tests/message-handler"
	"testing"
)

func TestHello(t *testing.T) {
	got := SayHello()
	want := "Hello, Go with test"

	if got!=want {
		errorMessage := fmt.Sprintf("got %s want %s\n",got, want);
		message.PrintError(t, &errorMessage);
	}
}