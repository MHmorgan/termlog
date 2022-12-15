Termlog
=======

Minimalistic, opinionated logging library for terminal
applications.

* Writes to `stderr` by default. May write to any `os.File`.
* Prints with colors and text formatting. This may be
  disabled, but not customized.
* Color and formatting is automatically disabled of the
  output file isn't a terminal.
* Each different type of log message has a different prefix,
  which cannot be changed.
* A timestamp may be added to the log messages. This is
  disabled by default.
* Thread-safe implementation.
* Any error returned by calls to `fmt` functions are
  ignored.
