package main

import(
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main(){
	fmt.Println("Server is now running")

	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from the web side")
	})
	server.Start(":3000")


}