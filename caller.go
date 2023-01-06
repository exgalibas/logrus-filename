/**
 * @author lin.tan
 * @date 2023/1/6
 * @description
 */

package logrus_filename

import (
	"runtime"
	"strings"
	"sync"
)

var once sync.Once
var logrusPackage string

const (
	maximumCallerDepth = 25
	minimumCallerDepth = 4
)

func getCaller(skip int) *runtime.Frame {

	once.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "fireHooks") {
				logrusPackage = getPackageName(funcName)
				break
			}
		}
	})

	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		if pkg != logrusPackage {
			if skip > 0 {
				skip--
				continue
			}
			return &f
		}
	}

	return nil
}

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}
