package main 

// go test executes the function that begins with Test
import "testing"

func TestName(t *testing.T) {
	name := getName()

	// report error if test fails 
	if name != "World!" {
		t.Error("Response from getName is unexpected value")
	}
}

