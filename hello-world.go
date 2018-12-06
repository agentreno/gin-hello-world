package main

import "os"
import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	message := os.Getenv("APP_MESSAGE")
	if len(message) == 0 {
		message = "pong"
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": message,
		})
	})

	r.GET("/file", func(c *gin.Context) {
        dat, err := ioutil.ReadFile("/mnt/store/output")
        if err != nil {
            dat := "error"
        }

		c.JSON(200, gin.H{
			"message": string(dat),
		})
	})
    return r
}

func main() {
    r := setupRouter()
	r.Run()
}
