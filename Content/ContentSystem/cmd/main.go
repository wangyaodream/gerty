package main

import (
	"fmt"
	"gerty/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.CmsRouter(r)

	err := r.Run()

	if err != nil {
		fmt.Printf("r run error = %v", err)
		return
	}

}
