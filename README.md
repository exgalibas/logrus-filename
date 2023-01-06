# logrus-filename
filename hook that support caller skip
see the issue [Need caller skip #972](https://github.com/sirupsen/logrus/issues/972)  
Although logrus does not support caller skip, we can still use hook to achieve it

# how to use
example

```go
// Package main in main.go
package main
import (
	"github.com/exgalibas/logrus-filename"
	"github.com/sirupsen/logrus"
)

const SkipKey = "@skip"

func main() {
	l := logrus.New()
	l.AddHook(logrus_filename.NewHook(logrus_filename.WithSkipKey(SkipKey)))
	TestSkip(l)
}
```
```go
package test

import "github.com/sirupsen/logrus"

func TestSkip(log *logrus.Logger) {
	log.Info("this is TestSkip")
	// output: INFO[0000] this is TestSkip file="/Users/exgalibas/gotest/main/test.go:52" 
	log.WithField(SkipKey, 1).Info("this is main")
	// output: INFO[0000] this is main file="/Users/exgalibas/gotest/main/main.go:36"
}
```
```go
//global hook
package main
import (
	"github.com/exgalibas/logrus-filename"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()
	l.AddHook(logrus_filename.NewHook(logrus_filename.WithSkip(1)))
	//now l always skip 1 caller
}
```
u can skip caller by logrus.WithField, and u can also define your SkipKey by the Options
```go
// options
WithSkip(int) -- define the caller skip depth directly, just for global
WithLogLevels([]logrus.Level) -- define which log levels trigger the hook
WithFormatter(HookFormatter) -- define file format in log, default func fileFormatter
WithSkipKey(string) -- skip callers at runtime, not global, just for once
WithRelease(bool) -- if true it will delete(entry.Data, SkipKey), default true
```
func `GetCaller` can help u to get `*runtime.Frame` with skip u specified