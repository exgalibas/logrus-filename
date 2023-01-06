/**
 * @author lin.tan
 * @date 2023/1/6
 * @description
 */

package logrus_filename

import "github.com/sirupsen/logrus"

type Option func(*Hook)

func WithSkip(skip int) Option {
	return func(h *Hook) {
		h.skipDepth = skip
	}
}

func WithLevels(levels []logrus.Level) Option {
	return func(h *Hook) {
		h.levels = levels
	}
}

func WithFormatter(formatter HookFormatter) Option {
	return func(h *Hook) {
		h.formatter = formatter
	}
}

func WithSkipKey(skipKey string) Option {
	return func(h *Hook) {
		h.skipKey = skipKey
	}
}

func WithRelease(release bool) Option {
	return func(h *Hook) {
		h.release = release
	}
}
