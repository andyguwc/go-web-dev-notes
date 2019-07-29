/*
type assertion canary tests that will fail quickly if you make a mistake on the interface definition 
alert basic failures in your assumptions 
*/

package canary

import (
	"io"
	"testing"
)

type MyWriter struct{}

func (m *MyWriter) Write([]byte) error {
	return nil
}

func main() {
	m := map[string]interface{}{
		"w": &MyWriter{},
	}
}

func doSomething(m map[string]interface{}) {
	w := m["w"].(io.Writer)
}

func TestWriter(t *testing.T) {
	var _ io.Writer = &MyWriter{}
}
