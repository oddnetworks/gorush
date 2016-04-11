package gopush

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

func PrintGoPushVersion() {
	fmt.Printf(`GoPush %s, Compiler: %s %s, Copyright (C) 2016 Bo-Yi Wu, Inc.`,
		Version,
		runtime.Compiler,
		runtime.Version())
}

func VersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		c.Writer.Header().Set("Server-Version", "GoPush/"+Version)
		c.Next()
	}
}