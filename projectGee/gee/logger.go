package gee

/*
middleware
*/

import (
	"log"
	"time"
)

// type HandlerFunc func(*Context)
func Logger() HandlerFunc {
	return func(c *Context) {
		// start timer
		t := time.Now()
		// process request
		c.Next()
		// calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
