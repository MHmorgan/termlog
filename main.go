package termlog

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/logrusorgru/aurora/v3"
	"github.com/mattn/go-isatty"
)

var (
	buf     []byte
	logfile *os.File
	mu      = &sync.Mutex{}

	au               aurora.Aurora
	colorsEnabled    = true
	timestampEnabled = false

	errorStyle   = func(s string) string { return au.Red(au.Bold(s)).String() }
	warningStyle = func(s string) string { return au.Yellow(au.Bold(s)).String() }
	goodStyle    = func(s string) string { return au.Green(au.Bold(s)).String() }
	infoStyle    = func(s string) string { return au.Faint(s).String() }
	emphStyle    = func(s string) string { return au.Bold(s).String() }

	errorPrefix   = "[!!]"
	warningPrefix = "[!]"
	goodPrefix    = "[✓]"
	badPrefix     = "[✗]"
	infoPrefix    = "[·]"
	emphPrefix    = "[*]"
)

func init() {
	SetLogFile(os.Stderr)
}

func useColor() bool {
	return colorsEnabled && isatty.IsTerminal(logfile.Fd())
}

// SetLogFile sets the file to write logs messages to.
// If the file is a TTY colors will be enabled, unless the
// user has explicitly disabled colors.
func SetLogFile(f *os.File) {
	logfile = f
	au = aurora.NewAurora(useColor())
}

// SetTimestampEnabled enables or disables timestamps in the log output.
// Timestamps are disabled by default.
func SetTimestampEnabled(b bool) {
	timestampEnabled = b
}

// SetColorsEnabled enables or disables colors (and bold formatting)
// in the log output.
// Colors are enabled by default.
func SetColorsEnabled(b bool) {
	colorsEnabled = b
	au = aurora.NewAurora(useColor())
}

func timestamp() string {
	if !timestampEnabled {
		return ""
	}
	s := time.Now().Format("15:04:05")
	if useColor() {
		s = au.Faint(s).String()
	}
	return s
}

// Output writes a log message to the log file. This is the
// most basic log function, and is used by the other log
// functions. May be used to write custom log messages.
//
// The style function is used to format the message (prefix
// and message). If style is nil, the message is not formatted.
//
// The prefix is the string that is printed before the message.
//
// If timestamp is enabled, the timestamp is printed between
// the prefix and the message. The timestamp styling is not
// customizable.
//
// A newline is added to the end of the message, unless the
// message already ends with a newline.
func Output(style func(string) string, prefix string, s string) {
	ts := timestamp()
	if style == nil {
		style = func(s string) string {
			return s
		}
	}
	mu.Lock()
	defer mu.Unlock()
	buf = buf[:0]
	buf = append(buf, style(prefix)...)
	buf = append(buf, ' ')
	if timestampEnabled {
		buf = append(buf, ts...)
		buf = append(buf, ' ')
	}
	buf = append(buf, style(s)...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}
	logfile.Write(buf)
}

// Print writes a message with custom prefix, without any
// formatting or coloring, to the log file.
func Print(prefix string, a ...any) {
	Output(nil, prefix, fmt.Sprint(a...))
}

// Printf is like Print, but takes a format string and arguments.
func Printf(prefix, format string, a ...any) {
	Output(nil, prefix, fmt.Sprintf(format, a...))
}

// Println is like Print.
func Println(prefix string, a ...any) {
	Output(nil, prefix, fmt.Sprintln(a...))
}

// Error writes an error message to the log file.
func Error(a ...any) {
	Output(errorStyle, errorPrefix, fmt.Sprint(a...))
}

// Errorf is like Error, but takes a format string and arguments.
func Errorf(format string, a ...any) {
	Output(errorStyle, errorPrefix, fmt.Sprintf(format, a...))
}

// Errorln is like Error.
func Errorln(a ...any) {
	Output(errorStyle, errorPrefix, fmt.Sprintln(a...))
}

// Fatal is equivalent to Error() followed by a call to os.Exit(1).
func Fatal(a ...any) {
	Output(errorStyle, errorPrefix, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf is equivalent to Errorf() followed by a call to os.Exit(1).
func Fatalf(format string, a ...any) {
	Output(errorStyle, errorPrefix, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalln is equivalent to Errorln() followed by a call to os.Exit(1).
func Fatalln(a ...any) {
	Output(errorStyle, errorPrefix, fmt.Sprintln(a...))
	os.Exit(1)
}

// FatalIfErr is equivalent to Fatal() if err is not nil,
// otherwise does nothing.
func FatalIfErr(err error) {
	if err != nil {
		Fatal(err)
	}
}

// Panic is equivalent to Error() followed by a call to panic().
func Panic(a ...any) {
	s := fmt.Sprint(a...)
	Output(errorStyle, errorPrefix, s)
	panic(s)
}

// Panicf is equivalent to Errorf() followed by a call to panic().
func Panicf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	Output(errorStyle, errorPrefix, s)
	panic(s)
}

// Panicln is equivalent to Errorln() followed by a call to panic().
func Panicln(a ...any) {
	s := fmt.Sprintln(a...)
	Output(errorStyle, errorPrefix, s)
	panic(s)
}

// PanicIfErr is equivalent to Panic() if err is not nil,
// otherwise does nothing.
func PanicIfErr(err error) {
	if err != nil {
		Panic(err)
	}
}

// Warn writes a warning message to the log file.
func Warn(a ...any) {
	Output(warningStyle, warningPrefix, fmt.Sprint(a...))
}

// Warnf is like Warn, but takes a format string and arguments.
func Warnf(format string, a ...any) {
	Output(warningStyle, warningPrefix, fmt.Sprintf(format, a...))
}

// Warnln is like Warn.
func Warnln(a ...any) {
	Output(warningStyle, warningPrefix, fmt.Sprintln(a...))
}

// Emph writes an emphasized message to the log file.
func Emph(a ...any) {
	Output(emphStyle, emphPrefix, fmt.Sprint(a...))
}

// Emphf is like Emph, but takes a format string and arguments.
func Emphf(format string, a ...any) {
	Output(emphStyle, emphPrefix, fmt.Sprintf(format, a...))
}

// Emphln is like Emph.
func Emphln(a ...any) {
	Output(emphStyle, emphPrefix, fmt.Sprintln(a...))
}

// Info writes an info message to the log file.
func Info(a ...any) {
	Output(infoStyle, infoPrefix, fmt.Sprint(a...))
}

// Infof is like Info, but takes a format string and arguments.
func Infof(format string, a ...any) {
	Output(infoStyle, infoPrefix, fmt.Sprintf(format, a...))
}

// Infoln is like Info.
func Infoln(a ...any) {
	Output(infoStyle, infoPrefix, fmt.Sprintln(a...))
}

// Good writes a success message to the log file.
func Good(a ...any) {
	Output(goodStyle, goodPrefix, fmt.Sprint(a...))
}

// Goodf is like Good, but takes a format string and arguments.
func Goodf(format string, a ...any) {
	Output(goodStyle, goodPrefix, fmt.Sprintf(format, a...))
}

// Goodln is like Good.
func Goodln(a ...any) {
	Output(goodStyle, goodPrefix, fmt.Sprintln(a...))
}

// Bad writes a failure message to the log file.
func Bad(a ...any) {
	Output(errorStyle, badPrefix, fmt.Sprint(a...))
}

// Badf is like Bad, but takes a format string and arguments.
func Badf(format string, a ...any) {
	Output(errorStyle, badPrefix, fmt.Sprintf(format, a...))
}

// Badln is like Bad.
func Badln(a ...any) {
	Output(errorStyle, badPrefix, fmt.Sprintln(a...))
}
