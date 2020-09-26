package encapsulated

import (
	"fmt"
	"io"
	"os"
)

type trace struct {
	noisy bool
	f     io.Writer
}

func (t trace) Print(s string) {
	if t.noisy {
		fmt.Fprintf(t.f, "%s\n", s)
	}
}

func (t *trace) On() {
	t.noisy = true
}

func (t *trace) Off() {
	t.noisy = false
}

type Trace interface {
	On()
	Off()
	Print(s string)
}

func TraceCustom(ff io.Writer) Trace {
	return &trace{f: ff}
}

func TraceDefault() Trace {
	return &trace{f: os.Stdout}
}
