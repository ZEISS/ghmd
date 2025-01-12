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
	return Write(w,
		Header(r.s.Header),
		Spacer(),
		Footer(r.s.Footer))
}

func New(s *spec.Spec) Renderer {
	return &renderer{s: s}
}

type Node interface {
	Render(w io.Writer) error
}

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

func Footer(footer string) Node {
	return &stringer{text: footer}
}

func Spacer() Node {
	return &stringer{text: "\n"}
}

func Write(out io.Writer, fn ...Renderer) error {
	for _, f := range fn {
		err := f.Render(out)
		if err != nil {
			return err
		}
	}

	return nil
}
