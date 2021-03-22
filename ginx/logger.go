package ginx

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ignoredPathPatterns = "/(liveness|healthy)$"
	jsonFormatter       = `{"time":"%s","client_ip":"%s","status":"%d","method":"%s","path":"%s","err_msg":"%s","latency":"%s"}` + "\n"
)

type Logger struct {
	ignoreRegexp *regexp.Regexp
}

// WithIgnoreLogPath 忽略指定路径的日志输出
// 例如: pattern="/(liveness|healthy)$"
func (l *Logger) WithIgnoreLogPath(pattern string) {
	l.ignoreRegexp = regexp.MustCompile(pattern)
}

func (l *Logger) formatter() func(param gin.LogFormatterParams) string {
	var defaultLoggerFormatter = func(param gin.LogFormatterParams) string {
		if l.ignoreRegexp.MatchString(param.Path) {
			return ""
		}

		if param.Latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			param.Latency = param.Latency - param.Latency%time.Second
		}

		return fmt.Sprintf(jsonFormatter,
			param.TimeStamp.Format("2006/01/02T15:04:05Z"),
			param.ClientIP,
			param.StatusCode,
			param.Method,
			param.Path,
			param.ErrorMessage,
			param.Latency,
		)
	}

	return defaultLoggerFormatter
}

// LogHandlerFunc 返回 gin 使用的 Log HandlerFunc
func (l *Logger) LogHandlerFunc() gin.HandlerFunc {
	fmtter := l.formatter()
	return gin.LoggerWithFormatter(fmtter)
}

// LogHandlerFuncWithIgnorePath 返回 gin 使用的 Log HandlerFunc
func LogHandlerFuncWithIgnorePath(patterns ...string) gin.HandlerFunc {
	l := Logger{}

	var patt string
	if len(patterns) == 0 {
		patt = ignoredPathPatterns
	} else {
		patt = fmt.Sprintf("(%s)", strings.Join(patterns, "|"))
	}
	l.WithIgnoreLogPath(patt)

	return l.LogHandlerFunc()
}
