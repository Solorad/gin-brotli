package main

import (
	brotli "github.com/Solorad/gin-brotli"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(brotli.Brotli(brotli.DefaultCompression))
	r.GET("/json", func(c *gin.Context) {
		bytes, err := ioutil.ReadFile("../testdata/data.json")
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, bytes)
	})
	r.Run()
}
