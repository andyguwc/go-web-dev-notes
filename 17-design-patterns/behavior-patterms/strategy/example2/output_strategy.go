package strategy

import "io"

type Output interface {
	Draw() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

// adding a LogWriter for logging purposes of outside packages 
type DrawOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func(d *DrawOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}

func(d *DrawOutput) SetWriter(w io.Writer) {
	d.Writer = w
}

