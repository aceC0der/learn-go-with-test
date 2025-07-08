package integers

import (
	"fmt"
	"learn-go-with-tests/message-handler"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		errorMessage := fmt.Sprintf("sum of 2+2 is '%d' expected '%d'", sum, expected)
		message.PrintErrorMessage(t, &errorMessage)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}