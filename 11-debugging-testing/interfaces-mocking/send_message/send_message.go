
package msg

type Message struct {
	// ...
}

func (m *Message) Send(email, subject string, body []byte) error {
	// ...
	return nil
}

// define an interface that describes the methods you use on Message
type Messager interface {
	Send(email, subject string, body []byte) error
}

func Alert(m Messager, problem []byte) error { // passes that interface intead of the message type
	return m.Send("noc@example.com", "Critical Error", problem)
}

