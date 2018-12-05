

# 算法实验报告

姓名：王宏斌     学号：U201514632     班级：ACM1501

## 实验1.匹配问题

### 目的

已知有若干医院和一定数量的学生，每个医院或学生均对对方有一个偏好排序，请扩展Gale–Shapley算法，使之支持：

1. 某些医院可接收多名学生 
2. 医院数量与学生数量不等 

### 要求

* 给出新算法的描述

* 编写该算法的程序
* 编制测试数据，给出实验结果：给出几个不同的数据集进行测试
* 分析、证明该匹配是否是稳定匹配

### 算法设计

#### 描述如下：

```
Initially all h∈H and s∈S are free
While there is a hospital h that has vacant place and hasn't proposed to every student
	Choose such a hospital h
	Let s be the highest-ranked student in h's preference list to whom h has not yet proposed
	If s is free student
		s is matched to h
	Else s is currently matched to h'
		If s prefers h to h' then
			s is matched to h
			h' vacates a place
		Endif
	Endif
EndWhile
Return the set of matched pairs
```

#### 正确性证明

假设存在一对医院 h、学生 s，使得 h 存在空闲位置或者 h 已满但是存在匹配的学生 s' 使得 h 相对于 s' 更加偏好 s，且 s 相对于当前匹配的医院更加偏好 h。

若 h 未满，则 h 对所有学生请求过了；若 h 已满，而 h 相对于 s' 更加偏好 s，则 h 必然在请求 s' 之前请求过 s 了，总之 h 请求过 s。最后 h、s 没有匹配，要么是已有更好的匹配被直接拒绝了，要么匹配之后有令 s 更加偏好的请求者从而将 h 淘汰，无论如何 s 当前的匹配者在 s 的偏好列表中一定优于 h，因此假设不成立。

综上该算法得出的匹配一定为稳定匹配

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
	"os"
)

var hospitalPrefs map[int][]int // every hospital's preference list
var studPrefs map[int][]int     // every student's preference list
var studRanks [][]int           // every student's rank in every hospital's preference list
var hospitalRanks [][]int       // every hospital's rank in every student's preference list
var studMatches []int           // studMatches[i] is the hospital that student i is matched to
var hospitalMatches []int       // hospitalMatches[i] is the lowest-rank student who is matched to hospital i

const INFILE = "matching.txt"

func input() {
	// hospitals[i] has default capicity i+1

	var n, m int
	hospitalPrefs = make(map[int][]int)
	studPrefs = make(map[int][]int)
	fmt.Scanf("%d%d", &n, &m)

	var tmp int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &tmp)
			hospitalPrefs[i] = append(hospitalPrefs[i], tmp)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &tmp)
			studPrefs[i] = append(studPrefs[i], tmp)
		}
	}
	fmt.Scanln()
}

func match() {
	numHospitals, numStuds := len(hospitalPrefs), len(studPrefs)
	hospitalMatchHist := make([]int, numHospitals)

	hospitalMatches = make([]int, numHospitals)
	studMatches = make([]int, numStuds)
	for stud := range studMatches {
		studMatches[stud] = -1
	}

	// get every hospital's rank in every student's preference list
	hospitalRanks = make([][]int, numStuds)
	for stud := range hospitalRanks {
		hospitalRanks[stud] = make([]int, numHospitals)
		for rank, hospital := range studPrefs[stud] {
			hospitalRanks[stud][hospital] = rank
		}
	}

	shouldContinue := true
	for shouldContinue {
		shouldContinue = false
		for hospital := range hospitalMatches {
			if hospitalMatches[hospital] > hospital || hospitalMatchHist[hospital] >= numStuds {
				// size >= capicaty or has proposed to every student
				continue
			}
			shouldContinue = true
			favoriate := hospitalPrefs[hospital][hospitalMatchHist[hospital]]
			hospitalMatchHist[hospital]++
			match := studMatches[favoriate]
			if match == -1 || hospitalRanks[favoriate][hospital] < hospitalRanks[favoriate][match] {
				studMatches[favoriate] = hospital
				hospitalMatches[hospital]++
			}
			if match != -1 {
				hospitalMatches[match]--
			}
		}
	}
}

func isStableMatch() bool {
	numHospitals, numStuds := len(hospitalPrefs), len(studPrefs)
	studRanks := make([][]int, numHospitals)
	for hospital := range studRanks {
		studRanks[hospital] = make([]int, numStuds)
		for rank, stud := range hospitalPrefs[hospital] {
			studRanks[hospital][stud] = rank
		}
	}

	worstMatchedStuds := make([]int, numHospitals)
	for hospital := range worstMatchedStuds {
		worstMatchedStuds[hospital] = -1
	}
	for stud, hospital := range studMatches {
		worstMatchedStud := worstMatchedStuds[hospital]
		if worstMatchedStud == -1 || studRanks[hospital][stud] > studRanks[hospital][worstMatchedStud] {
			worstMatchedStuds[hospital] = stud
		}
	}

	for hospital := 0; hospital < numHospitals; hospital++ {
		for stud := 0; stud < numStuds; stud++ {
			matchedStud, matchedHospital := worstMatchedStuds[hospital], studMatches[stud]
			if (matchedStud == -1 || studRanks[hospital][stud] < studRanks[hospital][matchedStud]) &&
				hospitalRanks[stud][hospital] < hospitalRanks[stud][matchedHospital] {
				fmt.Println(studMatches)
				fmt.Println("Unstable match", hospital, stud)
				return false
			}
		}
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
		input()
		match()
		fmt.Println(studMatches)
		if isStableMatch() {
			fmt.Println("OK!")
		}
	}
}
```

### 实验结果

#### 输入格式

* 第一行为一个正整数 k，表示有 k 组输入数据
* 每一组数据第一行两个正整数 m, n，表示有 m 个医院和 n 个学生，医院由整数编号 0~m 表示，学生由整数编号 0~n 表示，医院 i 可接收学生数量默认为 i+1
* 接下来 m 行每行 n 个整数，表示每个医院对学生的偏好排名，第 i 行第 j 列 s 代表医院 i-1 第 j+1 偏好的学生为 s
* 接下来 n 行每行 m 个整数，表示每个学生对医院的偏好排名，第 i 行第 j 列 h 代表学生 i-1 第 j+1 偏好的医院为 h

#### 输出格式

* 每组输出两行，第一行为各学生匹配到的医院编号，第 i 列 h 代表学生 i-1 分配到医院 h
* 第二行代表是否为稳定匹配，若是则输出 "OK!"，否则输出 "Unstable match" 及检测到的更优匹配

#### 输入文件 matching.txt：

```
3

5 10
9 8 7 6 5 4 3 2 1 0
6 3 8 4 1 2 5 7 0 9
7 6 5 4 0 1 2 3 9 8
4 3 5 6 9 0 1 2 8 7
0 8 1 4 6 2 3 5 7 9
4 3 2 1 0
0 1 2 3 4
3 2 0 1 4
3 2 1 0 4
0 2 3 4 1
0 3 2 1 4
2 0 3 1 4
1 3 2 0 4
2 1 4 0 3
4 3 2 1 0

4 10
9 3 8 4 2 1 5 7 0 6
5 6 7 1 0 4 2 8 9 3
5 6 3 4 9 0 8 2 1 7
0 8 2 6 4 1 3 5 7 9
3 2 1 0
0 1 2 3
3 2 0 1
3 0 1 2
0 2 3 1
0 3 2 1
2 0 3 1
1 3 2 0
2 1 0 3
3 2 1 0

4 9
4 2 8 3 1 5 7 0 6
5 6 7 4 2 0 1 8 3
5 6 3 4 0 7 1 2 8
6 2 8 0 4 1 3 5 7
3 2 1 0
0 1 2 3
3 2 0 1
3 2 1 0
0 2 3 1
0 3 2 1
2 0 3 1
1 3 2 0
2 1 0 3
```

#### 输出内容

```
[4 1 3 3 2 0 2 1 2 4]
OK!
[3 1 3 3 0 3 2 1 2 3]
OK!
[3 0 3 3 0 0 2 1 2]
OK!
```

### 结果分析

本次实验编写了测试函数 isStableMatch，遍历所有(医院, 学生)对，若存在一对医院 h、学生 s，使得 h 存在空闲位置或者 h 已满但是存在匹配的学生 s' 使得 h 相对于 s' 更加偏好 s，且 s 相对于当前匹配的医院更加偏好 h，则输出 "Unstable match: h s"，否则输出 "OK!"。

程序检测结果均为 "OK!"，表明该匹配确实为稳定匹配。



## 实验2.区间划分

### 目的

* 已知：n 个课程，课程 j 的开始时间为 sj，结束时间为 fj

* 目标：使用尽量少的教室，调度所有课堂，使没有两个课会在同一教室同一时间冲突。

### 要求

* 给出新算法的描述
* 编写该算法的程序
* 用给定及自行设计的测试数据测试程序，输出测试结果

### 算法设计

#### 描述如下：

```
Initially no classroom is allotted
Sort lectures by start time in ascending order
For lecture L in lectures
	If there exists an allotted classroom R whose lectures are all compatible with L then
		L is scheduled to R
	Else
		A new classroom is allotted for L
	Endif
Endfor
Return the classrooms with scheduled lectures
```

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
	"os"
	"sort"
)

const INFILE = "interval.txt"

type Lecture struct {
	start, finish float32
}

type ByStart []Lecture

func (l ByStart) Len() int           { return len(l) }
func (l ByStart) Less(i, j int) bool { return l[i].start < l[j].start }
func (l ByStart) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func schedule(lectures []Lecture) [][]Lecture {
	sort.Sort(ByStart(lectures))
	var rooms [][]Lecture
	for _, lecture := range lectures {
		flag := false
		for i, room := range rooms {
			if len(room) == 0 || room[len(room)-1].finish <= lecture.start {
				rooms[i] = append(room, lecture)
				flag = true
				break
			}
		}
		if !flag {
			rooms = append(rooms, []Lecture{lecture})
		}
	}
	return rooms
}

func input() []Lecture {
	var numLectures int
	fmt.Scanf("%d", &numLectures)
	lectures := make([]Lecture, numLectures)
	for i := 0; i < numLectures; i++ {
		fmt.Scanf("%f", &lectures[i].start)
	}
	for i := 0; i < numLectures; i++ {
		fmt.Scanf("%f", &lectures[i].finish)
	}
	fmt.Scanln()
	return lectures
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
		data := input()
		rsts := schedule(data)
		fmt.Printf("Test%d:\n", i+1)
		for i, rst := range rsts {
			fmt.Printf("%dth classroom: %v\n", i+1, rst)
		}
		fmt.Println()
	}
}

```

### 实验结果

#### 输入格式

- 第一行为一个正整数 k，表示有 k 组输入数据
- 每一组数据第一行一个正整数 n，表示共有 n 个 lecture
- 接下来两行每行 m 个整数，第一行每个数字表示各 lecture 的开始时间，第二行每个数字表示各 lecture 的结束时间

#### 输出格式

- 每组输出 n+1 行，第一行为 "Testn" 代表当前为 第 n 组
- 接下来每行为对应教室分配的 lectures 的开始/结束时间

#### 输入文件 interval.txt：

```
4

10
9 9 9 13 13 11 11 15 15 14
12.5 10.5 10.5 14.5 14.5 12.5 12.5 16.5 16.5 16.5

8
1 3 5 7 9 10 11 13
4 7 9 10 11 19 20 14

12
9 4 7 2 1 6 3 3 8 10 19 14
14 10 13 5 8 9 4 5 10 11 22 20

10
18 13 21 6 8 10 15 9 14 7
21 18 24 13 11 14 19 13 18 19
```

#### 输出内容

```
Test1:
1th classroom: [{9 12.5} {13 14.5} {15 16.5}]
2th classroom: [{9 10.5} {11 12.5} {13 14.5} {15 16.5}]
3th classroom: [{9 10.5} {11 12.5} {14 16.5}]

Test2:
1th classroom: [{1 4} {5 9} {9 11} {11 20}]
2th classroom: [{3 7} {7 10} {10 19}]
3th classroom: [{13 14}]

Test3:
1th classroom: [{1 8} {8 10} {10 11} {14 20}]
2th classroom: [{2 5} {6 9} {9 14} {19 22}]
3th classroom: [{3 4} {4 10}]
4th classroom: [{3 5} {7 13}]

Test4:
1th classroom: [{6 13} {13 18} {18 21} {21 24}]
2th classroom: [{7 19}]
3th classroom: [{8 11} {14 18}]
4th classroom: [{9 13} {15 19}]
5th classroom: [{10 14}]
```

### 结果分析

可以看出每组输出每间教室内的各 lecture 时间上没有冲突，且没有办法重新调度 lecture 使得使用的教室数量更少，因此结果正确。