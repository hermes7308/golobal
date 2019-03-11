package golobal

import (
	"flag"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Start() {
	// argument
	port := flag.Int("port", 7308, "port number")

	flag.Parse()

	// url
	e := echo.New()
	e.GET("/golobal/hash", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// start
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*port)))
}
