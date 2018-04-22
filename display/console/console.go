package console

import (
	"fmt"
	"io"
)

type Console struct {
	writer io.Writer
}

func (c *Console) Cls() error {
	// not implemented
	return nil
}

func (c *Console) Println(message string) error {
	_, err := fmt.Fprintln(c.writer, message)
	return err
}

func New(w io.Writer) (*Console, error) {
	return &Console{
		writer: w,
	}, nil
}
