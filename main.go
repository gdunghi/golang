package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func add(c echo.Context) error {
	var a, b int
	var err error
	if a, err = strconv.Atoi(c.Param("a")); err != nil {
		return c.String(http.StatusBadRequest, "invalid value a")
	}
	if b, err = strconv.Atoi(c.Param("b")); err != nil {
		return c.String(http.StatusBadRequest, "invalid value b")
	}

	return c.String(http.StatusOK, strconv.Itoa(a+b))
}

func caesar(c echo.Context) error {
	var s = c.Param("text")
	var rot int
	var err error
	if rot, err = strconv.Atoi(c.Param("rot")); err != nil {
		return c.String(http.StatusBadRequest, "invalid value a")
	}

	result := ""
	for _, v := range s {
		if v == 32 {
			result += string(32)
		} else {
			m := ((int(v) + int(rot)) % 26)
			result += string(m + 97)
		}
	}

	return c.String(http.StatusOK, result)
}

func callFizzBuzz(c echo.Context) error {

	var number int
	var err error
	if number, err = strconv.Atoi(c.Param("number")); err != nil {
		return c.String(http.StatusBadRequest, "invalid value number")
	}
	result := fizzBuzz(number)

	return c.String(http.StatusOK, result)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", hello)
	e.GET("/add/:a/:b", add)
	e.GET("/fizzbuzz/:number", callFizzBuzz)
	e.GET("/caesar/:text/:rot", caesar)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
