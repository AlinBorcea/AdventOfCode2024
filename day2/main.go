package main

import (
	"bufio"
	"fmt"
	"os"
)

type RecordStatus int

const (
	RecordOk RecordStatus = iota
	RecordLevelsIncrease
	RecordLevelsDecrease
	RecordLevelNoChange
	RecordLevelsUnstable
)

type Input struct {
	records [][]int
}

type Output struct {
	levelDifferences [][]int
	safeRecordsCount int
}

func main() {
	input := readInput("input.txt")

	output := Output{}
	//output.levelDifferences = calculateLevelDifferences(input.records)
	output.safeRecordsCount = calculateNumberOfSafeRecordsSimple(input.records)

	fmt.Println(output.safeRecordsCount)
}

func calculateNumberOfSafeRecordsSimple(records [][]int) int {
	count := 0

	for i := range records {
		increasingNumbers := 0
		decreasingNumbers := 0
		equalNumbers := 0
		badRanges := 0
		for j := 1; j < len(records[i]); j++ {
			if records[i][j-1] < records[i][j] { //increasing
				increasingNumbers++
			} else if records[i][j-1] > records[i][j] { //decreasing
				decreasingNumbers++
			} else if records[i][j-1] == records[i][j] { //equal
				equalNumbers++
			}

			x := records[i][j-1] - records[i][j]
			if x < 0 {
				x *= -1
			}
			if x < 1 || x > 3 {
				badRanges++
			}
		}

		mistakes := 0
		if increasingNumbers == len(records[i])-1 || decreasingNumbers == len(records[i])-1 {
			mistakes++
		}
		if badRanges > 0 {
			mistakes++
		}

		if mistakes <= 1 {
			count++
		}
	}

	return count
}

func calculateLevelDifferences(records [][]int) [][]int {
	var differencesOfLevels [][]int

	for i := range records {
		var differences []int
		for j := 1; j < len(records[i]); j++ {
			differences = append(differences, records[i][j-1]-records[i][j])
		}
		differencesOfLevels = append(differencesOfLevels, differences)
	}

	return differencesOfLevels
}

func calculateNumberOfSafeRecords(levelDifferences [][]int) int {
	count := 0

	for i := range levelDifferences {
		if recordStatus(levelDifferences[i]) == RecordOk {
			count++
		}
	}

	return count
}

func recordStatus(difs []int) RecordStatus {
	negatives := 0
	positives := 0
	badRange := 0
	zeros := 0

	for i := range difs {
		x := difs[i]
		if x == 0 {
			zeros++
		}
		if x < 0 {
			negatives++
		}
		if x > 0 {
			positives++
		}

		if x < 0 {
			x *= -1
		}

		if x > 3 {
			badRange++
		}
	}

	tolerance := 1
	if badRange > 0 {
		tolerance--
	}
	if zeros > 0 {
		tolerance--
	}

	if negatives == len(difs)-1 || positives == len(difs)-1 {
		tolerance--
	}

	if tolerance < 0 {
		return RecordLevelsUnstable
	}

	return RecordOk
}

func readInput(filename string) Input {
	input := Input{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		record := readRecordLine(line)
		input.records = append(input.records, record)
	}

	return input
}

func readRecordLine(line string) []int {
	var nums []int
	x := 0

	for i := range line {
		if line[i] < '0' || line[i] > '9' {
			nums = append(nums, x)
			x = 0
		} else {
			x = x*10 + int(line[i]-48)
		}
	}
	nums = append(nums, x)

	return nums
}
