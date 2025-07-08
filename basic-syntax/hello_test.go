package basic

import (
	"learn-go-with-tests/message-handler"
	"testing"
)

func TestHello(t *testing.T) {
	// group tests
	t.Run("Say hello to go by sending empty string", func(t *testing.T) {
		got := SayHello("")
		want := "Hello, Go with test"
	
		if got != want {
			message.PrintError(t, &got, &want);
		}
	})

	t.Run("Say hello to people by sending their name", func(t *testing.T){
		got := SayHello("Mridul")
		want := "Hello, Mridul"
	
		if got != want {
			message.PrintError(t, &got, &want);
		}
	})
}