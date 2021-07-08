package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/labstack/echo/engine/standard"
    "github.com/marmelab/snake-solver-server/lib"
)

func main() {
    type Data struct {
        Width int `json:"width"`
        Height int `json:"height"`
        Snake [][2]int `json:"snake"`
        Apple [2]int `json:"apple"`
    }

    type Response struct {
        Path []int
        PossibleMoves int
        ComputationTime int64
        BestMoveScore float32
        MaxTick int
    }

    e := echo.New()

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:9000", "http://0.0.0.0:9000"},
        AllowCredentials: true,
    }))

    e.POST("/", func(c echo.Context) error {
        d := new(Data)
        if err := c.Bind(d); err != nil {
            return err
        }

        path, possibleMoves, computationTime, bestMoveScore, maxTick := computer.GetPath(d.Width, d.Height, d.Snake, d.Apple)
        fmt.Println(path)
        return c.JSON(http.StatusOK, Response{path, possibleMoves, computationTime.Nanoseconds()/1e6, bestMoveScore, maxTick})
    })

    e.Run(standard.New(":1323"))
}
