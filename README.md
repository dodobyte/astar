# A* in Go
This is a very small (120 sloc) single file, single function A* implementation in Go.

I implemented it to use in games with tile-based maps. It's not trying to be super-generic or super-flexible. It's just small and straightforward.

The only function is `FindPath` and the only data structure is `Point`.

```go
type Point struct {
	X, Y int
}

func FindPath(grid [][]byte, start, end Point) []Point
```
