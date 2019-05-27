package web

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// Serve initializes and runs the web service
func Serve(c chan os.Signal, addr string, dataBase *gorm.DB) {
	r := gin.Default()

	r.POST("/open", handlePayload(dataBase))
	r.POST("/status", handlePayload(dataBase))
	r.POST("/close", handlePayload(dataBase))
	r.POST("/post", handlePayload(dataBase))
	r.POST("/api/post", handlePayload(dataBase))

	log.Println("Starting HTTP server on port:", addr)
	log.Fatal(r.Run(addr))
}
