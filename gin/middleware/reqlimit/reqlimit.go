package reqlimit

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Limit function is gin middleware to limit current requests
func ReqLimit(max int) gin.HandlerFunc {
	if max <= 0 {
		log.Panic("max must be more than 0")
	}
	semaphore := make(chan bool, max)

	return func(c *gin.Context) {
		select {
		case semaphore <- true:
			// Ok, managed to get a space in queue. execute the handler
			c.Next()
			<-semaphore // Don't forget to release a handle
		default:
			c.Status(http.StatusBadGateway)
		}
	}
}
