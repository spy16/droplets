package errors

import (
	"fmt"
	"io"
	"path"
	"runtime"
)

const depth = 32

type stack []frame

// Format formats the stack of Frames according to the fmt.Formatter interface.
//
//    %s	lists source files for each Frame in the stack
//    %v	lists the source file and line number for each Frame in the stack
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//    %+v   Prints filename, function, and line number for each Frame in the stack.
func (st stack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			for _, f := range st {
				fmt.Fprintf(s, "\n%+v", f)
			}
		case s.Flag('#'):
			fmt.Fprintf(s, "%#v", []frame(st))
		default:
			fmt.Fprintf(s, "%v", []frame(st))
		}
	case 's':
		fmt.Fprintf(s, "%s", []frame(st))
	}
}

type frame struct {
	fn   string
	file string
	line int
}

// Format formats the frame according to the fmt.Formatter interface.
//
//    %s    source file
//    %d    source line
//    %n    function name
//    %v    equivalent to %s:%d
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//    %+s   function name and path of source file relative to the compile time
//          GOPATH separated by \n\t (<funcname>\n\t<path>)
//    %+v   equivalent to %+s:%d
func (f frame) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		switch {
		case s.Flag('+'):
			fmt.Fprintf(s, "%s\n\t%s", f.fn, f.file)
		default:
			io.WriteString(s, path.Base(f.file))
		}
	case 'd':
		fmt.Fprintf(s, "%d", f.line)
	case 'n':
		io.WriteString(s, f.fn)
	case 'v':
		f.Format(s, 's')
		io.WriteString(s, ":")
		f.Format(s, 'd')
	}
}

func callStack(skip int) stack {
	var frames []frame
	var pcs [depth]uintptr

	n := runtime.Callers(skip, pcs[:])
	for i := 0; i < n; i++ {
		pc := pcs[i]
		frame := frameFromPC(pc)
		frames = append(frames, frame)
	}
	return stack(frames)
}

func frameFromPC(pc uintptr) frame {
	fr := frame{}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		fr.fn = "unknown"
	} else {
		fr.fn = fn.Name()
	}

	fr.file, fr.line = fn.FileLine(pc)
	return fr
}
