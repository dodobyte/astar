package astar

import (
	"math"
)

type Point struct {
	X, Y int
}

type node struct {
	gcost float64
	hcost float64
}

const crossCorner = false

func inside(g [][]byte, p Point) bool {
	return p.X >= 0 && p.X < len(g) && p.Y >= 0 && p.Y < len(g[0])
}

func cornerPass(g [][]byte, a, b Point) bool {
	if a.X == b.X || a.Y == b.Y {
		return false
	}
	p := Point{a.X, b.Y}
	q := Point{b.X, a.Y}
	return (inside(g, p) && g[p.X][p.Y] == 1) ||
		(inside(g, q) && g[q.X][q.Y] == 1)
}

func passable(g [][]byte, a, b Point) bool {
	if !inside(g, a) || g[a.X][a.Y] == 1 {
		return false
	}
	if crossCorner {
		return true
	}
	return !cornerPass(g, a, b)
}

func distance(a, b Point) float64 {
	dx := math.Abs(float64(a.X - b.X))
	dy := math.Abs(float64(a.Y - b.Y))
	if dx == 1 && dy == 1 {
		return math.Sqrt2
	}
	return dx + dy
}

func minCost(set map[Point]*node) Point {
	var pt Point
	min := math.MaxFloat64
	for p, n := range set {
		fcost := n.gcost + n.hcost
		switch {
		case fcost < min:
			min = fcost
			pt = p
		case fcost == min:
			if n.hcost < set[pt].hcost {
				min = fcost
				pt = p
			}
		}
	}
	return pt
}

func reverse(path []Point) []Point {
	i := 0
	j := len(path) - 1
	for i < j {
		path[i], path[j] = path[j], path[i]
		i++
		j--
	}
	return path
}

func FindPath(grid [][]byte, start, end Point) []Point {
	if grid[end.X][end.Y] == 1 {
		return nil
	}
	open := map[Point]*node{}
	closed := map[Point]*node{}
	parent := map[Point]Point{}
	open[start] = &node{0, distance(start, end)}
	for len(open) > 0 {
		cur := minCost(open)
		if cur == end {
			path := []Point{cur}
			for cur != start {
				cur = parent[cur]
				path = append(path, cur)
			}
			return reverse(path)
		}
		cnode := open[cur]
		delete(open, cur)
		closed[cur] = cnode
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				nbor := Point{cur.X + i, cur.Y + j}
				if _, ok := closed[nbor]; ok || !passable(grid, cur, nbor) {
					continue
				}
				gcost := cnode.gcost + distance(cur, nbor)
				if nnode, ok := open[nbor]; !ok {
					open[nbor] = &node{gcost, distance(nbor, end)}
				} else if gcost >= nnode.gcost {
					continue
				}
				parent[nbor] = cur
				open[nbor].gcost = gcost
			}
		}
	}
	return nil
}
