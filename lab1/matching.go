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

const INFILE = "/Users/coordinate36/program/algorithm/lab1/matching.txt"

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
