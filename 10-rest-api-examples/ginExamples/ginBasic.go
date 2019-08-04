/* Simple GIN framework application 
Return UTC time from Server

Gin allows you to do a lot of stuff with just a few lines of code; all the
boilerplate details are taken away. Coming to the preceding program, we are creating a
router with the gin.Default function. Then, we are attaching routes with REST verbs as
we did in go-restful; a route to the function handler. Then, we are calling the
Run function by passing the port to run. The default port will be 8080.
*/


package main

import (
	"time"
	
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pingTime", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})

	r.Run(":8000") // Default listen and serve on 0.0.0.0:8080
}
