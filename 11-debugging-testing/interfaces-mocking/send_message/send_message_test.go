
/*
create interfaces to describe the types you need to test
use those interfaces in code and write mock implementations for tests 
*/

package msg

import (
	"testing"
)


type MockMessage struct {
	email, subject string
	body           []byte
}

// MockMessage type implements Messager interface 
func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body
	return nil
}

func TestAlert(t *testing.T) {
	msgr := new(MockMessage)
	body := []byte("Critical Error")

	Alert(msgr, body) 

	// MockMessage provides the same functions that production code uses, but instead of sending the message it stores the data 
	if msgr.subject != "Critical Error" {
		t.Errorf("Expected 'Critical Error', Got '%s'", msgr.subject)
	}
	// ...
}