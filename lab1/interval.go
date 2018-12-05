package main

import (
	"fmt"
	"os"
	"sort"
)

const INFILE = "/Users/coordinate36/program/algorithm/lab1/interval.txt"

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
