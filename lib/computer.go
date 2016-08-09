package computer

import "sort"

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

func getBestPath(paths []path) path {
    sort.Sort(sort.Reverse(byScore(paths)))
    return paths[0]
}

func exploration(firstMove path, snake [][2]int, apple [2]int, size size, pathsChannel chan []path) {
    var paths []path
    paths = append(paths, firstMove)

    maxTick := 2
    for tick := 1; tick < maxTick; tick++ {
        for _, p := range paths {
            newSnake := moveSnake(snake, apple, p.Path)

            for _, possibleMove := range getPossibleMoves(size, newSnake) {
                newPath := append(p.Path, possibleMove)
                newScore := getMoveScore(size, possibleMove, newSnake, apple, tick)

                paths = append(paths, path{newPath, newScore})
            }
        }
    }

    pathsChannel <- paths
}

func GetPath(width int, height int, snake [][2]int, apple [2]int) []int {
    paths := make(chan []path)
    size := size{width, height}

    for _, possibleMove := range getPossibleMoves(size, snake) {
        firstMove := path{[]int{possibleMove}, getMoveScore(size, possibleMove, snake, apple, 1)}
        go exploration(firstMove, snake, apple, size, paths);
    }

    p := <- paths

    return getBestPath(p).Path
}
