package shapes

import "github.com/sayden/go-design-patterns/behavioral/strategy/example2"

type TextSquare struct {
	strategy.DrawOutput
}

// So now the Draw() method is slightly different because, instead of writing directly to the
// console by using the println function, we have to write whichever io.Writer is stored in
// the Writer field

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}


// when we use t.Writer, we are actually accessing PrintOutput.Writer. The TextSquare type has a Writer field because
// the PrintOutput struct has it and it's embedded on the TextSquare struct.

