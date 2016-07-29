package computer

import "sort"

const maxTick = 10
const up, right, down, left = 0, 1, 2, 3
const block = 1;

type path struct {
    Path []int
    Score float32
}

type byScore []path

func (a byScore) Len() int { return len(a) }
func (a byScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

func isSnakeEatApple(snake [][2]int, apple [2]int) bool {
    return isSnakeHeadAtPosition(snake, apple)
}

func moveSnake(snake [][2]int, apple [2]int, moves []int) [][2]int {
    for _, move := range moves {
        snakeHead := getSnakeHead(snake)
        nextPosition := getAdjacentPosition(snakeHead, move)

        if !isSnakeEatApple(snake, apple) {
            snake = snake[1:]
        }

        snake = append(snake, nextPosition)
    }

    return snake
}

func getAdjacentPosition(position [2]int, move int) [2]int {
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
    return getSnakeHead(snake) == position
}

func isOutsideBoundingBox(position [2]int, grid [][]int) bool {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    var width = len(grid[0])
    var height = len(grid)

    if x >= width || x < 0 || y >= height || y < 0 {
        return true
    }

    return false
}

func isEmptyCell(grid [][]int, position [2]int) bool {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    if grid[x][y] != block {
        return true
    }

    return false
}

func getPossibleMoves(grid [][]int, snake [][2]int) []int {
    var snakeHead = getSnakeHead(snake)
    var snakeTail = getSnakeTail(snake)

    var possibleMoves []int
    for _, move := range []int{up, right, down, left} {
        var adjacentPosition = getAdjacentPosition(snakeHead, move)

        if (!isOutsideBoundingBox(adjacentPosition, grid) && isEmptyCell(grid, adjacentPosition)) || adjacentPosition == snakeTail {
            possibleMoves = append(possibleMoves, move)
        }
	}

    return possibleMoves
}

func getSnakeHead(snake [][2]int) [2]int {
    return snake[len(snake) - 1]
}

func getSnakeTail(snake [][2]int) [2]int {
    return snake[0]
}

func initializeGrid(width int, height int, snake [][2]int, apple [2]int) [][]int {
    var grid [][]int

    for y := 0; y < height; y++ {
		grid = append(grid, []int{})
		for x := 0; x < width; x++ {
			grid[y] = append(grid[y], 0)
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

func isSnakeHasFreeSpace(grid [][]int, snake [][2]int) bool {
    return len(getPossibleMoves(grid, snake)) > 0
}

func getMoveScore(grid [][]int, move int, snake [][2]int, apple [2]int, tick int) float32 {
    newSnake := moveSnake(snake, apple, []int{move})

    if isSnakeHeadAtPosition(newSnake, apple) {
        if !isSnakeHasFreeSpace(grid, newSnake) {
            return float32(0)
        }

        return (float32(1) / float32(tick)) * float32(100)
    }

    return float32(1)
}

func getBestPath(paths [][]int, scores []float32) path {
    var pathsSelected []path
    for index, score := range scores {
        pathsSelected = append(pathsSelected, path{paths[index], score})
    }

    sort.Sort(sort.Reverse(byScore(pathsSelected)))
    return pathsSelected[0]
}

func GetPath(width int, height int, snake [][2]int, apple [2]int) []int {
    var paths [][]int
    var scores []float32

    grid := initializeGrid(width, height, snake, apple)

    for _, possibleMove := range getPossibleMoves(grid, snake) {
        paths = append(paths, []int{possibleMove})
        scores = append(scores, getMoveScore(grid, possibleMove, snake, apple, 1))
    }

    for tick := 1; tick < maxTick; tick++ {
        var newPaths [][]int
        var newScores []float32

        for index, path := range paths {
            newSnake := moveSnake(snake, apple, path)
            grid = initializeGrid(width, height, newSnake, apple)

            for _, possibleMove := range getPossibleMoves(grid, newSnake) {
                newPath := append(path, possibleMove)
                newPaths = append(newPaths, newPath)
                newScore := getMoveScore(grid, possibleMove, newSnake, apple, tick)

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
