package public

import (
	"fmt"
	"io"
)

type Trace struct {
	noisy bool
	F     io.Writer
}

func (t Trace) Print(s string) {
	if t.noisy {
		_, _ = fmt.Fprintf(t.F, "%s\n", s)
	}
}

func (t *Trace) On() {
	t.noisy = true
}

func (t *Trace) Off() {
	t.noisy = false
}
