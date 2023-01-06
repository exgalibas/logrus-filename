/**
 * @author lin.tan
 * @date 2023/1/6
 * @description
 */

package logrus_filename

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type HookFormatter func(*Hook, *logrus.Entry) error

type Hook struct {
	skipDepth int
	skipKey   string
	levels    []logrus.Level
	formatter HookFormatter
	release   bool
}

func (hook *Hook) Levels() []logrus.Level {
	return hook.levels
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	if hook.skipKey != "" {
		if skipValue, ok := entry.Data[hook.skipKey]; ok {
			if skipInt, ok := skipValue.(int); ok {
				hook.skipDepth = skipInt
			}
			if hook.release {
				delete(entry.Data, hook.skipKey)
			}
		}
	}
	return hook.formatter(hook, entry)
}

func NewHook(options ...Option) *Hook {
	hook := &Hook{
		formatter: fileFormatter,
		release:   true,
	}

	for _, option := range options {
		option(hook)
	}

	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}

	return hook
}

func fileFormatter(hook *Hook, entry *logrus.Entry) error {
	f := getCaller(hook.skipDepth)
	if f != nil {
		entry.Data["file"] = fmt.Sprintf("%s:%d", f.File, f.Line)
	}

	return nil
}
