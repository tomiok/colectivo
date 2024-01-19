package colectivo

import (
	"io"
	"os"
)

type Subscriber interface {
	Handle(e Event) error
}

type StdOutSub struct {
	W io.Writer
}

func (s *StdOutSub) Handle(e Event) error {
	_, err := s.W.Write(e.Bytes)
	return err
}

func WithStdOut() *StdOutSub {
	return &StdOutSub{W: os.Stdout}
}
