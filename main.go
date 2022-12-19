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
	au               aurora.Aurora
	logfile          *os.File
	mtx              = &sync.Mutex{}
	timestampEnabled = false
	colorsEnabled    = true
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
	s := fmt.Sprintf("%s ", time.Now().Format("15:04:05"))
	if useColor() {
		s = au.Faint(s).String()
	}
	return s
}

// Log writes a log message to the log file. This is the
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
// The message may be formatted as described in `fmt`.
func Log(style func(string) string, prefix string, a ...any) {
	if style == nil {
		style = func(s string) string {
			return s
		}
	}

	msg := style(fmt.Sprint(a...))
	prefix = style(prefix)
	if timestampEnabled {
		prefix = fmt.Sprintf("%s %s", prefix, timestamp())
	}

	mtx.Lock()
	defer mtx.Unlock()
	_, _ = fmt.Fprintln(logfile, prefix, msg)
}

// Logf is like Log, but takes a format string and arguments.
func Logf(style func(string) string, prefix, format string, a ...interface{}) {
	Log(style, prefix, fmt.Sprintf(format, a...))
}

// Err writes an error message to the log file.
func Err(a ...interface{}) {
	style := func(s string) string {
		return au.Red(au.Bold(s)).String()
	}
	Log(style, "[!!]", a...)
}

// Errf is like Err, but takes a format string and arguments.
func Errf(format string, a ...interface{}) {
	Err(fmt.Sprintf(format, a...))
}

// Fatal is like Err, but exits the program with exit code 1.
func Fatal(a ...interface{}) {
	Err(a...)
	os.Exit(1)
}

// Fatalf is like Fatal, but takes a format string and arguments.
func Fatalf(format string, a ...interface{}) {
	Errf(format, a...)
	os.Exit(1)
}

// Warn writes a warning message to the log file.
func Warn(a ...interface{}) {
	style := func(s string) string {
		return au.Yellow(au.Bold(s)).String()
	}
	Log(style, "[!]", a...)
}

// Warnf is like Warn, but takes a format string and arguments.
func Warnf(format string, a ...interface{}) {
	Warn(fmt.Sprintf(format, a...))
}

// Emph writes an emphasized message to the log file.
func Emph(a ...interface{}) {
	style := func(s string) string {
		return au.Bold(s).String()
	}
	Log(style, "[*]", a...)
}

// Emphf is like Emph, but takes a format string and arguments.
func Emphf(format string, a ...interface{}) {
	Emph(fmt.Sprintf(format, a...))
}

// Info writes an info message to the log file.
func Info(a ...interface{}) {
	style := func(s string) string {
		return au.Faint(s).String()
	}
	Log(style, "[·]", a...)
}

// Infof is like Info, but takes a format string and arguments.
func Infof(format string, a ...interface{}) {
	Info(fmt.Sprintf(format, a...))
}

// Good writes a success message to the log file.
func Good(a ...interface{}) {
	style := func(s string) string {
		return au.Green(au.Bold(s)).String()
	}
	Log(style, "[✓]", a...)
}

// Goodf is like Good, but takes a format string and arguments.
func Goodf(format string, a ...interface{}) {
	Good(fmt.Sprintf(format, a...))
}

// Bad writes a failure message to the log file.
func Bad(a ...interface{}) {
	style := func(s string) string {
		return au.Red(au.Bold(s)).String()
	}
	Log(style, "[✗]", a...)
}

// Badf is like Bad, but takes a format string and arguments.
func Badf(format string, a ...interface{}) {
	Bad(fmt.Sprintf(format, a...))
}
