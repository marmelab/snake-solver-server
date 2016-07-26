package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/marmelab/snake-solver-server/lib"
)

const width, height = 5, 5

func main() {
    grid := [width][height]int{
        {1, 1, 1, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 2},
    }

    snake := [][2]int{
        {0, 0},
        {0, 1},
        {0, 2},
    }

    apple := [2]int{4, 4}

    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        path := lib.GetPath(grid, snake, apple)
        return c.JSON(http.StatusOK, path)
    })
    e.Run(standard.New(":1323"))
}
