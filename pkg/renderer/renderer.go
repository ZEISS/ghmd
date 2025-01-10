package renderer

import (
	"io"

	"github.com/zeiss/ghmd/pkg/spec"
)

var _ Renderer = (*renderer)(nil)

type Renderer interface {
	Render(w io.Writer) error
}

type renderer struct {
	s *spec.Spec
}

func (r *renderer) Render(w io.Writer) error {
	if err := Header(r.s.Header).Render(w); err != nil {
		return err
	}

	return nil
}

func New(s *spec.Spec) Renderer {
	return &renderer{s: s}
}

type Node interface {
	Render(w io.Writer) error
}

type statefulWriter struct {
	w   io.Writer
	err error
}

// Write is a node that writes to the stateful writer.
func (w *statefulWriter) Write(p []byte) {
	if w.err != nil {
		return
	}
	_, w.err = w.w.Write(p)
}

var _ Node = (*stringer)(nil)

type stringer struct {
	text string
}

func (s *stringer) Render(w io.Writer) error {
	_, err := w.Write([]byte(s.text))
	return err
}

func Header(header string) Node {
	return &stringer{text: header}
}
