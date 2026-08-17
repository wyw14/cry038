package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"sync/atomic"
)

var seq uint64

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-ID")
		if id == "" {
			id = "class-" + strconv.FormatUint(atomic.AddUint64(&seq, 1), 10)
		}
		c.Set("request_id", id)
		c.Header("X-Request-ID", id)
		c.Header("Referrer-Policy", "no-referrer")
		c.Next()
	}
}
