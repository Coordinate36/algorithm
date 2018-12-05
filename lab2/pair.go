package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct{ X, Y int }
type ByX []Point
type ByY []Point

type Pair struct {
	A, B     Point
	Distance int
}

const INFILE = "/Users/coordinate36/program/algorithm/lab2/pair.txt"

func (p ByX) Len() int           { return len(p) }
func (p ByX) Less(i, j int) bool { return p[i].X < p[j].X }
func (p ByX) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p ByY) Len() int           { return len(p) }
func (p ByY) Less(i, j int) bool { return p[i].X < p[j].X }
func (p ByY) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func input() []Point {
	var n int
	fmt.Scanf("%d", &n)
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d%d", &points[i].X, &points[i].Y)
	}
	fmt.Scanln()
	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(p1 Point, p2 Point) int {
	return (p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)
}

func bruteforce(points []Point, n int) (pair Pair) {
	pair.Distance = math.MaxInt32
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := distance(points[i], points[j])
			if d < pair.Distance {
				pair.Distance = d
				pair.A, pair.B = points[i], points[j]
			}
		}
	}
	return
}

func stripClosest(strip []Point, minD int) (pair Pair) {
	for i := 0; i < len(strip); i++ {
		for j := i + 1; j < len(strip) && strip[j].Y-strip[j].Y < minD; j++ {
			d := distance(strip[i], strip[j])
			if d < minD {
				minD = d
				pair.A, pair.B = strip[i], strip[j]
			}
		}
	}
	pair.Distance = minD
	return
}

func closestPairRec(PX, PY []Point, n int) Pair {
	if n <= 3 {
		return bruteforce(PX, n)
	}

	mid := n >> 1
	LX, RX := PX[:mid], PX[mid:]

	var LY, RY []Point
	midPoint := PX[mid]
	for _, point := range PY {
		if point.X < midPoint.X {
			LY = append(LY, point)
		} else {
			RY = append(RY, point)
		}
	}

	lClosest := closestPairRec(LX, LY, mid)
	rClosest := closestPairRec(RX, RY, n-mid)
	var best Pair
	if lClosest.Distance < rClosest.Distance {
		best = lClosest
	} else {
		best = rClosest
	}

	var strip []Point
	for _, point := range PY {
		if abs(point.X-midPoint.X) < best.Distance {
			strip = append(strip, point)
		}
	}
	crossClosest := stripClosest(strip, best.Distance)
	if crossClosest.Distance < best.Distance {
		best = crossClosest
	}

	return best
}

func closestPair(points []Point) Pair {
	PX := make([]Point, len(points))
	PY := make([]Point, len(points))
	copy(PX, points)
	copy(PY, points)

	sort.Sort(ByX(PX))
	sort.Sort(ByY(PY))
	return closestPairRec(PX, PY, len(points))
}

func isClosestPair(points []Point, pair Pair) bool {
	best := bruteforce(points, len(points))
	if best.Distance < pair.Distance {
		return false
	}
	return true
}

func main() {

	// redirect stdin
	inFile, err := os.Open(INFILE)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()
	os.Stdin = inFile

	var numTests int
	fmt.Scanf("%d", &numTests)
	fmt.Scanln()

	for i := 0; i < numTests; i++ {
		points := input()
		pair := closestPair(points)
		fmt.Printf("Closest point pair: %v %v. Distance: %v\n", pair.A, pair.B, math.Sqrt(float64(pair.Distance)))
		if isClosestPair(points, pair) {
			fmt.Println("OK!")
		} else {
			fmt.Println("Error!")
		}
	}
}
