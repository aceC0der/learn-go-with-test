package iteration

import (
	"learn-go-with-tests/message-handler"
	"testing"
)

func TestRepeat(t *testing.T){
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected  {
		message.PrintError(t, &repeated, &expected)
	}
}