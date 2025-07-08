package basic

import (
	"learn-go-with-tests/message-handler"
	"testing"
)

func TestHello(t *testing.T) {
	// group tests
	t.Run("Say hello to go by sending empty string in english", func(t *testing.T) {
		got := SayHello("", "English")
		want := "Hello, Go with test"
	
		if got != want {
			message.PrintError(t, &got, &want);
		}
	})

	t.Run("Say hello to people by sending their name in english", func(t *testing.T){
		got := SayHello("Mridul", "English")
		want := "Hello, Mridul"
	
		if got != want {
			message.PrintError(t, &got, &want);
		}
	})

	t.Run("Say hello to go by sending empty string in spanish", func(t *testing.T) {
		got := SayHello("", "Spanish")
		want := "Hola, Go with test"
	
		if got != want {
			message.PrintError(t, &got, &want);
		}
	})

	t.Run("Say hello to people by sending their name in english", func(t *testing.T){
		got := SayHello("Mridul", "Spanish")
		want := "Hola, Mridul"
	
		if got != want {
			message.PrintError(t, &got, &want);
		}
	})
}