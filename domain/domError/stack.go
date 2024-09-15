package domError

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type (
	_stackframe struct {
		counter     uintptr
		packageName string
		fnName      string
		file        string
		line        int
	}
)

func newStackframe(c uintptr) _stackframe {

	stack := _stackframe{counter: c}

	rFn := runtime.FuncForPC(stack.counter)
	if rFn == nil {
		return stack
	}

	stack.file, stack.line = rFn.FileLine(stack.counter - 1)

	name := rFn.Name()
	_package := ""
	// find last / on name
	if index := strings.LastIndex(name, "/"); index >= 0 {
		_package = name[:index] + "/"
		name = name[index+1:]
	}

	// split pacakge and function name
	if index := strings.Index(name, "."); index >= 0 {
		_package = name[:index]
		name = name[index+1:]
	}

	stack.fnName = name
	stack.packageName = _package

	return stack
}

func (f _stackframe) toString() string {
	str := fmt.Sprintf("%s:%d \n", f.file, f.line)

	var source string

	if f.line > 0 {
		file, err := os.Open(f.file)
		if err == nil {
			defer file.Close()

			scanner := bufio.NewScanner(file)
			currentLine := 1
			for scanner.Scan() {
				if currentLine == f.line {
					source = string(bytes.Trim(scanner.Bytes(), " \t"))
					break
				}
				currentLine++
			}
			if err := scanner.Err(); err != nil {
				source = ""
			}
		}
	}

	if source != "" {
		str = str + fmt.Sprintf("\t%s=> %s\n", f.fnName, source)
	}

	return str
}
