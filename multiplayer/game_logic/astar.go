package game_logic

import (
	// "log"
	// "strconv"
	"math"
)

// type Map struct {
// 	TileSize int           `json:"tailsize"`
// 	SizeX    int           `json:"sizex"`
// 	SizeY    int           `json:"sizey"`
// 	Field    [5][5]int	   `json:"field"`
// }

type Point struct {
	XCell int
	YCell int

	G int // g(x). Стоимость пути от начальной вершины. У start g(x) = 0
	F int // f(x) = g(x) + h(x)
}

func (p *Point) IsEqual(other *Point) bool {
	if p == nil || other == nil {
		return true
	}
	if p.XCell == other.XCell && p.YCell == other.YCell {
		return true
	}
	return false
}

type Points []*Point

func (points *Points) IsExists(p *Point) (result bool) {
	result = false
	for i := 0; i < len(*points); i++ {
		if (*points)[i].IsEqual(p) {
			result = true
			break
		}
	}
	return
}

func (points *Points) Remove(p *Point) {
	for i := 0; i < len(*points); i++ {
		if (*points)[i].IsEqual(p) {
			*points = append((*points)[:i], (*points)[i+1:]...)
			break
		}
	}
}

func (points *Points) Add(p *Point) {
	*points = append(*points, p)
}

func minF(points Points) (p *Point) {
	if len(points) < 1 {
		return nil
	}
	p = points[0]
	for i := 1; i < len(points); i++ {
		if points[i].F < p.F {
			p = points[i]
		}
	}
	return
}

func getUnclosedNeighbours(p *Point, m *Map, close *Points) Points {
	var result Points
	if p.XCell > 0 {
		var left = &Point{
			XCell: p.XCell - 1,
			YCell: p.YCell,
		}
		if m.Field[left.YCell][left.XCell] == 0 && !close.IsExists(left) {
			result = append(result, left)
		}
	}
	if p.XCell < len(m.Field)-1 {
		var right = &Point{
			XCell: p.XCell + 1,
			YCell: p.YCell,
		}
		if m.Field[right.YCell][right.XCell] == 0 && !close.IsExists(right) {
			result = append(result, right)
		}
	}
	if p.YCell > 0 {
		var bottom = &Point{
			XCell: p.XCell,
			YCell: p.YCell - 1,
		}
		if m.Field[bottom.YCell][bottom.XCell] == 0 && !close.IsExists(bottom) {
			result = append(result, bottom)
		}
	}
	if p.YCell < len(m.Field[p.XCell])-1 {
		var top = &Point{
			XCell: p.XCell,
			YCell: p.YCell + 1,
		}
		if m.Field[top.YCell][top.XCell] == 0 && !close.IsExists(top) {
			result = append(result, top)
		}
	}
	return result
}

func absDist(first *Point, second *Point) int {
	return int(0.5 + math.Sqrt(math.Pow(float64(first.XCell-second.XCell), 2)+math.Pow(float64(first.YCell-second.YCell), 2)))
}

func getWay(from map[*Point]*Point, start *Point, curr *Point) Points {
	var points Points
	for !from[curr].IsEqual(start) {
		points = append(points, curr)
		curr = from[curr]
	}
	points = append(points, curr)
	points = append(points, start)
	return points
}

func heuristics(curr *Point, goal *Point) int {
	// 0.5 + - Костыль для округления (не робит при отриц, так что все ок)
	return int(0.5 + math.Sqrt(math.Pow(float64(curr.XCell-goal.XCell), 2)+math.Pow(float64(curr.YCell-goal.YCell), 2)))
}

//func AStar(start *Point, goal *Point, m *Map) (Points, bool) {
//	if goal.IsEqual(start) {
//		var points Points
//		points = append(points, start)
//		return points, true
//	}
//
//	var open Points
//	var closed Points
//	var from = make(map[*Point]*Point)
//	start.G = 0
//	start.F = start.G + heuristics(start, goal)
//
//	open = append(open, start)
//	for len(open) > 0 {
//		curr := minF(open)
//		if curr.IsEqual(goal) {
//			return getWay(from, start, curr), true
//		}
//		open.Remove(curr)
//		closed.Add(curr)
//		unclosedNeighbours := getUnclosedNeighbours(curr, m, &closed)
//		for _, neighbour := range unclosedNeighbours {
//			var tempG = curr.G + absDist(curr, neighbour)
//			if !open.IsExists(neighbour) || tempG < neighbour.G {
//				from[neighbour] = curr
//				neighbour.G = tempG
//				neighbour.F = neighbour.G + heuristics(neighbour, goal)
//			}
//			if !open.IsExists(neighbour) {
//				open.Add(neighbour)
//			}
//		}
//	}
//	return nil, false
//}
