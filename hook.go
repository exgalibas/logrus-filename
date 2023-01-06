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
	SkipDepth int
	SkipKey   string
	LogLevels []logrus.Level
	Formatter HookFormatter
	Release   bool
}

func (hook *Hook) Levels() []logrus.Level {
	return hook.LogLevels
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	if hook.SkipKey != "" {
		if skipValue, ok := entry.Data[hook.SkipKey]; ok {
			if skipInt, ok := skipValue.(int); ok {
				hook.SkipDepth = skipInt
			}
			if hook.Release {
				delete(entry.Data, hook.SkipKey)
			}
		}
	}
	return hook.Formatter(hook, entry)
}

func NewHook(options ...Option) *Hook {
	hook := &Hook{
		Formatter: fileFormatter,
		Release:   true,
	}

	for _, option := range options {
		option(hook)
	}

	if len(hook.LogLevels) == 0 {
		hook.LogLevels = logrus.AllLevels
	}

	return hook
}

func fileFormatter(hook *Hook, entry *logrus.Entry) error {
	f := GetCaller(hook.SkipDepth)
	if f != nil {
		entry.Data["file"] = fmt.Sprintf("%s:%d", f.File, f.Line)
	}

	return nil
}
