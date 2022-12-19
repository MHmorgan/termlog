Termlog
=======

An opinionated logging library for terminal applications.

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
* Any errors encountered when writing to the log file
  is ignored.


Example
-------

```go
package main

import (
	tl "github.com/mhmorgan/termlog"
)

func main() {
	tl.Info("Hello, world!")
	
	tl.SetTimestampEnabled(true)
	tl.Error("Something went wrong!")
}
```
