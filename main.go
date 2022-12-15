package termlog

import (
	"fmt"
	"os"
	"time"

	"github.com/logrusorgru/aurora/v3"
	"github.com/mattn/go-isatty"
)

var (
	au               aurora.Aurora
	logfile          *os.File
	logger           = myLogger{messages: make(chan string)}
	timestampEnabled = false
	colorsEnabled    = true
)

func init() {
	SetLogFile(os.Stderr)
	go logger.run()
}

type myLogger struct {
	messages chan string
}

func (l myLogger) run() {
	for msg := range l.messages {
		_, _ = fmt.Fprintln(logfile, msg)
	}
}

func (l myLogger) log(msg string) {
	l.messages <- msg
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
func Log(style func(string) string, prefix, msg string, args ...interface{}) {
	if style == nil {
		style = func(s string) string {
			return s
		}
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	if timestampEnabled {
		msg = fmt.Sprintf("%s %s %s", style(prefix), timestamp(), style(msg))
	} else {
		msg = fmt.Sprintf("%s %s", style(prefix), style(msg))
	}

	logger.log(msg)
}

// Err writes an error message to the log file.
func Err(msg string, args ...interface{}) {
	style := func(s string) string {
		return au.Red(au.Bold(s)).String()
	}
	Log(style, "[!!]", msg, args...)
}

// Warn writes a warning message to the log file.
func Warn(msg string, args ...interface{}) {
	style := func(s string) string {
		return au.Yellow(au.Bold(s)).String()
	}
	Log(style, "[!]", msg, args...)
}

// Emph writes an emphasized message to the log file.
func Emph(msg string, args ...interface{}) {
	style := func(s string) string {
		return au.Bold(s).String()
	}
	Log(style, "[*]", msg, args...)
}

// Info writes an info message to the log file.
func Info(msg string, args ...interface{}) {
	style := func(s string) string {
		return au.Faint(s).String()
	}
	Log(style, "[·]", msg, args...)
}

// Good writes a success message to the log file.
func Good(msg string, args ...interface{}) {
	style := func(s string) string {
		return au.Green(au.Bold(s)).String()
	}
	Log(style, "[✓]", msg, args...)
}

// Bad writes a failure message to the log file.
func Bad(msg string, args ...interface{}) {
	style := func(s string) string {
		return au.Red(au.Bold(s)).String()
	}
	Log(style, "[✗]", msg, args...)
}
