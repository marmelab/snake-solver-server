package computer

import "sort"
import "time"

const up, right, down, left = 0, 1, 2, 3
const block = 1;

type path struct {
    Path []int
    Score float32
}

type size struct {
    width int
    height int
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

        snake = append(snake, nextPosition)

        if !isSnakeEatApple(snake, apple) {
            snake = snake[1:]
        }
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

func isOutsideBoundingBox(size size, position [2]int) bool {
    var x = position[0] // @FIXME: use destructuring
    var y = position[1]

    if x >= size.width || x < 0 || y >= size.height || y < 0 {
        return true
    }

    return false
}

func isEmptyCell(position [2]int, snake [][2]int) bool {
    for _, s := range snake {
        if s == position {
            return false
        }
    }

    return true
}

func getPossibleMoves(size size, snake [][2]int) []int {
    var snakeHead = getSnakeHead(snake)
    var snakeTail = getSnakeTail(snake)

    var possibleMoves []int
    for _, move := range []int{up, right, down, left} {
        var adjacentPosition = getAdjacentPosition(snakeHead, move)

        if (!isOutsideBoundingBox(size, adjacentPosition) && isEmptyCell(adjacentPosition, snake)) || adjacentPosition == snakeTail {
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

func isSnakeHasFreeSpace(size size, snake [][2]int) bool {
    return len(getPossibleMoves(size, snake)) > 0
}

func isLastMove(size size, snake [][2]int) bool {
    if len(snake) == (size.width * size.height) - 1 {
        return true
    }

    return false
}

func getMoveScore(size size, move int, snake [][2]int, apple [2]int, tick int) float32 {
    newSnake := moveSnake(snake, apple, []int{move})

    if isSnakeHeadAtPosition(newSnake, apple) {
        if !isSnakeHasFreeSpace(size, newSnake) && !isLastMove(size, snake) {
            return float32(0)
        }

        return (float32(1) / float32(tick)) * float32(100)
    }

    return float32(1)
}

func getBestPath(paths [][]int, scores []float32) []int {
    var pathsSelected []path
    for index, score := range scores {
        pathsSelected = append(pathsSelected, path{paths[index], score})
    }

    sort.Sort(sort.Reverse(byScore(pathsSelected)))
    return pathsSelected[0].Path
}

func GetPath(width int, height int, snake [][2]int, apple [2]int) []int {
    var paths [][]int
    var scores []float32

    size := size{width, height}

    for _, possibleMove := range getPossibleMoves(size, snake) {
        paths = append(paths, []int{possibleMove})
        scores = append(scores, getMoveScore(size, possibleMove, snake, apple, 1))
    }

    if isLastMove(size, snake) {
        return getBestPath(paths, scores)
    }

    var totalTime time.Duration
    tick := 1
    for {
        startTimeTick := time.Now()

        var newPaths [][]int
        var newScores []float32

        for index, path := range paths {
            newSnake := moveSnake(snake, apple, path)

            for _, possibleMove := range getPossibleMoves(size, newSnake) {
                newPath := append(path, possibleMove)
                newPaths = append(newPaths, newPath)
                newScore := getMoveScore(size, possibleMove, newSnake, apple, tick)

                if newScore > scores[index] {
                    newScores = append(newScores, newScore)
                } else {
                    newScores = append(newScores, scores[index])
                }
            }
        }

        if len(newScores) > 0 && len(newPaths) > 0 {
            paths = newPaths
            scores = newScores
        }

        elapsedTimeTick := time.Since(startTimeTick)
        totalTime += elapsedTimeTick
        tick++

        if totalTime >= 1 * time.Second {
            return getBestPath(paths, scores)
        }
    }

    return getBestPath(paths, scores)
}
