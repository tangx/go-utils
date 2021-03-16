package ginx

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ignoredPathPatterns = "[/liveness|/healthy]$"
	ignoredLogRegexp    = regexp.MustCompile(ignoredPathPatterns)
)

// GinLogger is the default log format function Logger middleware uses.
var defaultLoggerFormatter = func(param gin.LogFormatterParams) string {
	if ignoredLogRegexp.MatchString(param.Path) {
		return ""
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	jsonFormatter := `{"time":"%s","client_ip":"%s","status":"%d","method":"%s","path":"%s","err_msg":"%s","latency":"%s"}`
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

func defaultLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(defaultLoggerFormatter)
}
