package domError

import (
	"bytes"
	"fmt"
)

type Error struct {
	Err    error
	stack  []uintptr
	frames []_stackframe
}

func (err *Error) Error() string {

	msg := err.Err.Error()

	return msg
}

func (err *Error) Stack() string {
	buf := bytes.Buffer{}

	err.frames = make([]_stackframe, len(err.stack))

	for i, pc := range err.stack {
		err.frames[i] = newStackframe(pc)
	}

	for _, f := range err.frames {
		buf.WriteString(f.toString())
	}

	return fmt.Sprintf("domError: %s, with stack trace: \n\n%s\n", err.Error(), string(buf.Bytes()))

}
