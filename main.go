package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// LuckyColor aaa
type LuckyColor struct {
	BirthMonth int    `json:"birthMonth"`
	LuckyColor string `json:"luckyColor"`
}

// About aaa
type About struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	License     string `json:"license"`
	CreatedAt   string `json:"createdAt"`
}

// Error aaa
type Error struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	// Routing
	e.GET("/", func(c echo.Context) error {
		path := "./README.md"
		readme, err := ioutil.ReadFile(path)
		if err != nil {
			c.String(http.StatusOK, "Could't Open README!")
		}
		return c.String(http.StatusOK, string(readme))
	})

	e.GET("/about", func(c echo.Context) error {
		about := About{
			Title:       "Golang-API-Server",
			Description: "This repository is a template of API-Server with Golang.",
			Author:      "Taku Kitamura",
			License:     "MIT",
			CreatedAt:   "2018-03-11 16:28:57",
		}
		return c.JSON(http.StatusOK, about)
	})

	e.GET("/luckyColor/:birthMonth", func(c echo.Context) error {
		colors := []string{
			"white",
			"black",
			"blue",
			"red",
			"pink",
			"yellow",
			"green",
		}

		birthMonth, err := strconv.Atoi(c.Param("birthMonth"))

		if err != nil {
			fmt.Println("Should be appropriate param.")
			return c.JSON(http.StatusBadRequest,
				Error{
					Message: "Should be appropriate param.",
				},
			)
		}

		if birthMonth < 1 || birthMonth > 12 {
			fmt.Println("Should be appropriate birthMonth.")
			return c.JSON(http.StatusBadRequest,
				Error{
					Message: "Should be appropriate birthMonth.",
				},
			)
		}

		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(colors))
		luckyColor := LuckyColor{
			BirthMonth: birthMonth,
			LuckyColor: colors[i],
		}

		return c.JSON(http.StatusOK, luckyColor)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
