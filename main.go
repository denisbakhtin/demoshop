package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/denisbakhtin/demoshop/controllers"
	"github.com/denisbakhtin/demoshop/system"
	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	mode := flag.String("mode", "debug", "Application mode: debug, release, test")
	flag.Parse()

	system.SetMode(mode)
	system.Init()

	gin.SetMode(system.GetMode())
	router := gin.Default()
	router.StaticFS("/public", http.Dir("public"))
	router.GET("/api/products", controllers.ProductsGet)
	router.GET("/api/products/:id", controllers.ProductGet)

	log.Fatal(router.Run(":7990"))
}
