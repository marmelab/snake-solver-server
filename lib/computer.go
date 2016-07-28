package computer

import "sort"

const maxTick = 6
const width, height = 5, 5
const up, right, down, left = 0, 1, 2, 3
const block = 1;
const apple = 2;

type Path struct {
    Path []int
    Score float32
}

type ByScore []Path

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

func MoveSnake(snake [][2]int, moves []int) [][2]int {

    for _, move := range moves {
        snakeHead := GetSnakeHead(snake)
        nextPosition := GetAdjacentPosition(snakeHead, move)

        snake = snake[1:]
        snake = append(snake, nextPosition)
    }

    return snake
}

func GetAdjacentPosition(position [2]int, move int) [2]int {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    switch move {
    case up:
        return [2]int{x - 1, y}
    case right:
        return [2]int{x, y + 1}
    case down:
        return [2]int{x + 1, y}
    case left:
        return [2]int{x, y - 1}
    default:
        return [2]int{x, y + 1} // @FIXME: return false
    }
}

func isSnakeHeadAtPosition(snake [][2]int, position [2]int) bool {
    return GetSnakeHead(snake) == position
}

func IsOutsideBoundingBox(position [2]int, grid [width][height]int) bool {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    var width = len(grid[0])
    var height = len(grid)

    if x >= width || x < 0 || y >= height || y < 0 {
        return true
    }

    return false
}

func IsEmptyCell(grid [width][height]int, position [2]int) bool {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    if grid[x][y] != block {
        return true
    }

    return false
}

func GetPossibleMoves(grid [width][height]int, snake [][2]int) []int {
    var head = GetSnakeHead(snake)

    var possibleMoves []int
    for _, move := range []int{up, right, down, left} {
        var adjacentPosition = GetAdjacentPosition(head, move)

        if !IsOutsideBoundingBox(adjacentPosition, grid) && IsEmptyCell(grid, adjacentPosition) {
            possibleMoves = append(possibleMoves, move)
        }
	}

    return possibleMoves
}

func GetSnakeHead(snake [][2]int) [2]int {
    return snake[len(snake) - 1]
}

func InitializeGrid(snake [][2]int, apple [2]int) [width][height]int {
    var grid [width][height]int
    var x, y int
    for x = 0; x < width; x++ {
        for y = 0; y < height; y++ {
            grid[x][y] = 0
        }
    }

    for _, snakePosition := range snake {
        xSnakePosition := snakePosition[0]
        ySnakePosition := snakePosition[1]

        grid[xSnakePosition][ySnakePosition] = block;
    }

    xApple := apple[0]
    yApple := apple[1]
    grid[xApple][yApple] = 2

    return grid
}

func getMoveScore(move int, snake [][2]int, apple [2]int, tick int) float32 {
    newSnake := MoveSnake(snake, []int{move})

    if isSnakeHeadAtPosition(newSnake, apple) {
        return (float32(1) / float32(tick)) * float32(100)
    }

    return float32(1)
}

func getBestPath(paths [][]int, scores []float32) Path {
    var pathsSelected []Path
    for index, score := range scores {
        pathsSelected = append(pathsSelected, Path{paths[index], score})
    }

    sort.Sort(sort.Reverse(ByScore(pathsSelected)))
    return pathsSelected[0]
}

func GetPath(grid [width][height]int, snake [][2]int, apple [2]int) []int {
    var paths [][]int
    var scores []float32

    for _, possibleMove := range GetPossibleMoves(grid, snake) {
        paths = append(paths, []int{possibleMove})
        scores = append(scores, getMoveScore(possibleMove, snake, apple, 1))
    }

    for tick := 1; tick < maxTick; tick++ {
        var newPaths [][]int
        var newScores []float32

        for index, path := range paths {
            newSnake := MoveSnake(snake, path)
            grid = InitializeGrid(newSnake, apple)

            for _, possibleMove := range GetPossibleMoves(grid, newSnake) {
                newPath := append(path, possibleMove)
                newPaths = append(newPaths, newPath)
                newScore := getMoveScore(possibleMove, newSnake, apple, tick)

                if newScore > scores[index] {
                    newScores = append(newScores, newScore)
                } else {
                    newScores = append(newScores, scores[index])
                }
            }
        }

        paths = newPaths
        scores = newScores
    }

    return getBestPath(paths, scores).Path
}
