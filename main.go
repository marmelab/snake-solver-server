package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/labstack/echo/engine/standard"
    "github.com/marmelab/snake-solver-server/lib"
)

const width, height = 5, 5

func main() {
    type Data struct {
        // Width int `json:"width" xml:"width" form:"width"` // @TODO
        // Height int `json:"height" xml:"height" form:"height"` // @TODO
        Snake [][2]int `json:"snake" xml:"snake" form:"snake"`
        Apple [2]int `json:"apple" xml:"apple" form:"apple"`
    }

    e := echo.New()

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:9000", "http://localhost"},
        AllowCredentials: true,
    }))

    e.POST("/", func(c echo.Context) error {
        d := new(Data)
        if err := c.Bind(d); err != nil {
            return err
        }

        grid := computer.InitializeGrid(d.Snake, d.Apple)
        path := computer.GetPath(grid, d.Snake, d.Apple)

        return c.JSON(http.StatusOK, path)
    })

    e.Run(standard.New(":1323"))
}
