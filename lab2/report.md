

# 算法实验报告

姓名：王宏斌     学号：U201514632     班级：ACM1501

## 实验1.字符串比对

### 目的

已知两个字符串 x~1~...x~n~, y~1~...y~m~; 比对成本为

1. c(i, j): x~i~ to y~j~ (Equals 0 if match, otherwise > 0) 
2. a(i): leaving out x~i~(>0)
3. b(j): leaving out y~j~(>0)

请给出最小成本的一个比对

### 要求

- 给出算法的两种动态规划描述
- 编写其中一种算法的程序
- 编制测试数据，给出实验结果：给出几个不同的数据集进行测试

### 算法设计

#### 状态转移方程

OPT(i, j) = min[c(x~i~, y~j~) + OPT(i-1, j-1), a(x~i~) + OPT(i-1, j), b(y~j~) + OPT(i, j-1)]	(1)

#### Basic DP：

> Alignment(X, Y)
> ​	Array A[0...m, 0...n]
> ​	Initialize A[i,0] = iδ for each i
> ​	Initialize A[0,j] = jδ for each j
> ​	For j = 1,...,n
> ​		For i = 1,...,m
> ​			Use state transition equation (1) to compute A[i, j]
> ​		Endfor
> ​	Endfor
> ​	Return A[m,n]

#### Space-Efficient DP:

> Alignment(X, Y)
> ​	Array B[0...m, 0...1]
> ​	Initialize B[i,0] = iδ for each i
> ​	For j = 1,...,n
> ​		B[0,1] = jδ
> ​		For i = 1,...,m
> ​			B[i,1] = min[c(x~i~, y~j~) + B(i-1, 0), a(x~i~) + B(i-1, 1), b(y~j~) + B(i, 0)]
> ​		Endfor
> ​		Move column 1 of B to column 0 to make room for next iteration
> ​	Endfor
> ​	Return B[m,1]

### 实验环境

系统：macOS

内核：Darwin Kernel Version 18.2.0

编程语言：go

编译环境：go1.10.3 darwin/amd64

### 实验过程

#### Space-Efficient DP:

```go
package main

import (
	"fmt"
	"os"
)

const INFILE = "alignment.txt"

var len1, len2 int
var str1, str2 []int
var gapPenality []int
var mismatchPenality [][]int

func input() {
	var m int
	fmt.Scanf("%d", &m)
	gapPenality = make([]int, m)
	mismatchPenality = make([][]int, m)
	for i := range mismatchPenality {
		mismatchPenality[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &gapPenality[i])
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &mismatchPenality[i][j])
		}
	}

	fmt.Scanf("%d", &len1)
	str1 = make([]int, len1)
	for i := 0; i < len1; i++ {
		fmt.Scanf("%d", &str1[i])
	}
	fmt.Scanf("%d", &len2)
	str2 = make([]int, len2)
	for i := 0; i < len2; i++ {
		fmt.Scanf("%d", &str2[i])
	}

	fmt.Scanln()
}

func min(first int, args ...int) int {
	for _, v := range args {
		if first > v {
			first = v
		}
	}
	return first
}

func align() int {
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len2; i++ {
		dp[0][i] = dp[0][i-1] + gapPenality[str2[i-1]]
	}
	for i := 1; i <= len1; i++ {
		dp[1][0] += gapPenality[str1[i-1]]
		for j := 1; j <= len2; j++ {
			dp[1][j] = min(dp[1][j-1]+gapPenality[str2[j-1]], dp[0][j-1]+mismatchPenality[str1[i-1]][str2[j-1]], dp[0][j]+gapPenality[str1[i-1]])
		}
		for j := 0; j <= len2; j++ {
			dp[0][j] = dp[1][j]
		}
	}

	return dp[1][len2]
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
		input()
		minPenality := align()
		fmt.Println("minimum penality:", minPenality)
	}
}
```

### 实验结果

#### 输入格式

- 第一行为一个正整数 k，表示有 k 组输入数据
- 每一组数据第一行一个正整数 m，表示字符集的大小
- 接下来一行 m 个非负整数，其中第 i 个表示 δ(i)
- 接下来一行表示第一个字符串 s1，首先一个正整数 n 表示长度，然后 n 个非负整数表示第一个字符串内容
- 一组数据的最后一行表示第二个字符串 s2，格式同上

#### 输出格式

- 每组数据输出一行 "minimum penality: W"，其中 W 为最小penalty

#### 提示

* 相对原题目，这里使用了非负整数来代替字符
* 上述 i 和 j 均从0开始计数，即满足 0 <= i, j < m
* 对于给定 i 和 j，不保证 α(i, j) = α(j, i)

#### 输入文件 alignment.txt：

```
3

2
1 1
0 1
1 0
3 0 0 1
3 1 0 0

4
1 2 3 4
0 1 1 2
1 0 2 2
1 2 0 1
2 2 1 0
4 0 0 1 2
4 0 1 2 3

4
1 1 1 1
0 1 1 2
1 0 2 2
1 2 0 1
2 2 1 0
4 0 0 1 2
4 0 1 2 3
```

#### 输出内容

```
minimum penality: 2
minimum penality: 4
minimum penality: 2
```

### 结果分析

可以看出三组输入的最佳对齐方式分别为:

```
0 0 1
1 0 0
```

```
0 0 1 2
0 1 2 3
```

```
0 0 1 2 -
0 - 1 2 3
```

最小 penality 分别为 2、4、2，从而程序输出结果正确。

## 实验2.最近点对

### 目的

- 已知：二维平面上的 n 个点，坐标 (x~i~, y~i~)
- 目标：输出欧式距离最近的点对及其对应距离

### 要求

- 给出算法描述，以及算法的合理性分析
- 编写该算法的程序
- 用自行设计的测试数据测试程序，输出测试结果

### 算法设计

#### 描述如下：

> ClosestPair(P)
> ​	Construct Px(sort P by x-coordinate) and Py(sort P by y-coordinate)
> ​	Return ClosestPairRec(Px, Py)
>
> ClosestPairRec(Px, Py)
> ​	If |P| <= 3 then
> ​		find closest pair by measuring all pairwise distances
> ​	Endif
> ​	
> ​	Divide Px into two equally sized part Qx、Rx by centermost x-coordinate
> ​	Divide Py into two equally sized part Qy、Ry by centermost x-coordinate
> ​	(q~0~, q~1~) = ClosestPairRec(Qx, Qy)
> ​	(r~0~, r~1~) = ClosestPairRec(Rx, Ry)
> ​	
> ​	δ = min(distance(q~0~, q~1~), distance(r~0~, r~1~))
> ​	x^*^ = maximum x-coordinate of a point in set Qx
> ​	L = {(x, y) | x = x^*^}
> ​	S = points in P within distance δ of L.
> ​	Sy = S sorted by y-coordinate in ascending order
> ​	
> ​	Construct Sy from Py
> ​	For each point s in Sy
> ​		Candidates = points whose y-coordinate is between (s.y, δ+s.y)
> ​		compute distance from s to each point in Candidates
> ​    	Let s, s' be pair achieving minimum of these distances
> ​    
> ​    	If d(s, s') < δ then
> ​    		Return (s, s')
> ​    	Else if d(q~0~, q~1~) < d(r~0~, r~1~)
> ​    		Return (q~0~, q~1~)
> ​    	Else
> ​    		Return (r~0~, r~1~)

#### 正确性分析

当 |P| <= 3 时，通过计算所有点对距离得到最近点对，显然正确

对于给定点集 P，将 P 分为两个几乎等大(点数相等或相差 1)的点集 Q、R，假设最近点对为 s、s'，则要么 s、s' 都来自于 Q，要么都来自于 R，要么 s 来自于 Q 而 s' 来自于 R(反之亦然)，于是问题分解为规模更小的问题，即求解 Q、R 的最近点对，然后对结果做归并。这样不断递归问题规模会越来越小。

已知当问题规模足够小时，即 |P| <= 3 时算法正确，从而只要归并过程无误即可。

归并过程：令 (q~0~, q~1~) 、(r~0~, r~1~) 分别为子问题 Q、R 中的最近点对，令 δ = min(distance(q~0~, q~1~), distance(r~0~, r~1~))。假设分别来自于 Q、R 中的点对 (s, s') 为 P 中的最近点对，则 distance(s, s') < δ => s.x - s'.x < distance(s, s') < δ and s.y - s'.y < distance(s, s') < δ，从而只用找 Q、R x-coordinate 分界线左右 δ 范围内的点作为 s、s' 的候选集 S，且对于 S 中的每个点 p 只用比较 y-coordinate 上下 δ 范围内的点到自身的距离，这里预先对 S 按 y-coordinate 升序排列，方便选择 S 中 y-coordinate 在 (p.y, δ + p.y) 范围内的点。可知归并过程理论上是合理的，且可通过对 S 的平面划分分析得知归并过程的复杂度为 O(Cn)，其中 C 为常数。

从而算法在理论上正确，且复杂度为

W(n) = 2W(n/2) + Cn = O(Cnlogn)

### 实验环境

系统：macOS

内核：Darwin Kernel Version 18.2.0

编程语言：go

编译环境：go1.10.3 darwin/amd64

### 实验过程

```go
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
```

### 实验结果

#### 输入格式

- 第一行为一个正整数 k，表示有 k 组输入数据
- 每一组数据第一行一个正整数 n，表示共有 n 个 点
- 接下来 n 行每行 2 个整数，分别代表各点的 X、Y 坐标

#### 输出格式

- 每组输出 2 行，第一行为 "Closest point pair: {x~0~ y~0~} {x~1~ y~1~}. Distance: D"。其中 {x~0~ y~0~} {x~1~ y~1~} 代表最近点对，D 代表该点对之间的距离
- 接下来一行表示检测最近点对正确性的结果，若找到更近点对则输出 "Error!"，否则输出 "OK!"

#### 输入文件 pair.txt：

```
3

6
2 3
12 30
40 50
5 1
12 10
3 4

8
1 2
2 5
4 3
8 10
9 1
5 2
7 9
6 1

9
0 0
1 3
5 4
3 4
8 9
7 6
8 10
6 6
100 100
```

#### 输出内容

```
Closest point pair: {2 3} {3 4}. Distance: 1.4142135623730951
OK!
Closest point pair: {7 9} {8 10}. Distance: 1.4142135623730951
OK!
Closest point pair: {8 9} {8 10}. Distance: 1
OK!
```

### 结果分析

本次实验编写了测试函数 isClosestPair，遍历所有点对，若存在点对之间的距离小于程序找出最近点对之间的距离则输出 "Error!"，否则输出 "OK!"。

程序检测结果均为 "OK!"，表明程序输出结果正确。