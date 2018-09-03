package main

import(
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main(){
	// setting up the router to the default one used by gin
	router := gin.Default()
	router.Use(static.Serve("/", static.Localfile("./views", true)))
	

	router.run(":3000")

}
