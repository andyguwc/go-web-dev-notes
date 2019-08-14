
/*

Adapter Pattern

Fit two incompatible interfaces 
1. Create an Adapter object that implements the ModernPrinter interface.
2. The new Adapter object must contain an instance of the LegacyPrinter
interface.
3. When using ModernPrinter, it must call the LegacyPrinter interface under
the hood, prefixing it with the text Adapter.

*/

package structural

import "fmt"

//------------------------------------------------------------------------

type LegacyPrinter interface {
	Print(s string) string
}

//------------------------------------------------------------------------

type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg)
	return
}

//------------------------------------------------------------------------

type NewPrinter interface {
	PrintStored() string
}

// new adapter contain an instance of legacy instance 
type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

// new adapter object implements the modern printer interface
func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}

	return
}

//------------------------------------------------------------------------
