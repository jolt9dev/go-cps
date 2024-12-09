package cps

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/jolt9dev/go-platform"
)

const (
	ARCH     = runtime.GOARCH
	PLATFORM = runtime.GOOS
)

var (
	Args     = os.Args[1:]
	Stderr   = os.Stderr
	Stdout   = os.Stdout
	Stdin    = os.Stdin
	ExecPath = os.Args[0]
	history  = []string{}
	reader   *bufio.Reader
	writer   *bufio.Writer
	eol      = []byte(platform.EOL)
)

func init() {
	current, _ := Cwd()
	history = append(history, current)
}

// Getwd returns a rooted path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// Getwd may return any one of them.
func Cwd() (string, error) {
	return os.Getwd()
}

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.
//
// For portability, the status code should be in the range [0, 125].
func Exit(code int) {
	os.Exit(code)
}

// Getpid returns the process id of the caller.
func Pid() int {
	return os.Getpid()
}

// Getpid returns the parent process id of the caller.
func Ppid() int {
	ppid := os.Getppid()
	if ppid == 0 {
		return -1
	}

	return ppid
}

// Getuid returns the numeric user id of the caller.
//
// On Windows, it returns -1.
func Uid() int {
	return os.Getuid()
}

// Getgid returns the numeric group id of the caller.
//
// On Windows, it returns -1.
func Gid() int {
	return os.Getgid()
}

// Geteuid returns the numeric effective user id of the caller.
//
// On Windows, it returns -1.
func Euid() int {
	return os.Geteuid()
}

// Getegid returns the numeric effective group id of the caller.
//
// On Windows, it returns -1
func Egid() int {
	return os.Getegid()
}

// Pushd changes the current working directory to the given path
// and adds the path to the directory stack.
//
// Parameters:
//   - path: the path to change to
func Pushd(path string) error {
	history = append(history, path)
	return os.Chdir(path)
}

// Popd changes the current working directory to the last path
func Popd() error {
	if len(history) == 1 {
		return nil
	}

	last := history[len(history)-1]
	history = history[:len(history)-1]
	return os.Chdir(last)
}

// Read reads data from the standard input.
func Read(b []byte) (int, error) {
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}

	return reader.Read(b)
}

// ReadLine reads a line from the standard input.
func ReadLine() (string, error) {
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}

	b, _, e := reader.ReadLine()
	return string(b), e
}

// ReadRune reads a single Unicode code point from the standard input.
func WriteBytes(b []byte) (int, error) {
	return Stdout.Write(b)
}

// WriteRune writes a single Unicode code point to the standard output.
func WriteRune(r rune) (int, error) {
	if writer == nil {
		writer = bufio.NewWriter(os.Stdout)
	}

	return writer.WriteRune(r)
}

// WriteBytes writes the bytes to the standard output.
func WriteString(s string) (int, error) {
	if writer == nil {
		writer = bufio.NewWriter(Stdout)
	}

	b, err := writer.WriteString(s)
	if err != nil {
		return b, err
	}
	writer.Flush()
	return b, err
}

// Writef writes the formatted string to the standard output.
func Writef(format string, a ...interface{}) (int, error) {
	msg := fmt.Sprintf(format, a...)
	return WriteString(msg)
}

// Writeln writes the string to the standard output and appends a newline.
func Writeln(s string) (int, error) {

	n, err := WriteString(s)
	if err != nil {
		return n, err
	}

	n2, err := writer.Write(eol)
	return n + n2, err
}
