package golobal

import (
	"encoding/json"
	"flag"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Start() {
	// argument
	port := flag.Int("PORT", 7309, "port number")

	flag.Parse()

	// url
	e := echo.New()
	e.GET("/golobal", func(c echo.Context) error {
		url := c.QueryParams()["url"][0]
		hashInfo := ExtractHashInfo(url)

		result, err := json.Marshal(hashInfo)
		if err != nil {
			return c.String(http.StatusOK, err.Error())
		}

		return c.String(http.StatusOK, string(result))
	})

	// start
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*port)))
}
