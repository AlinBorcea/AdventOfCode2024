package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Input struct {
	numbers1 []int
	numbers2 []int
}

type Output struct {
	numberDistances []int
	distancesSum    int
	similarityScore int
}

func main() {
	input := readInput("input.txt")

	output := Output{}
	output.numberDistances = calculateNumberDistances(input)
	output.distancesSum = calculateDistancesSum(output.numberDistances)
	output.similarityScore = calculateSimilarityScore(input)

	fmt.Printf("Value of distancesSum is %v\n", output.distancesSum)
	fmt.Printf("Value of similarityScore is %v\n", output.similarityScore)
}

func calculateNumberDistances(input Input) []int {
	sort.Slice(input.numbers1, func(i, j int) bool { return input.numbers1[i] < input.numbers1[j] })
	sort.Slice(input.numbers2, func(i, j int) bool { return input.numbers2[i] < input.numbers2[j] })

	var distances []int
	for i := 0; i < len(input.numbers1); i++ {
		dif := input.numbers1[i] - input.numbers2[i]
		if dif < 0 {
			dif *= -1
		}

		distances = append(distances, dif)
	}

	return distances
}

func calculateSimilarityScore(input Input) int {
	score := 0

	for i := range input.numbers1 {
		appearances := 0
		for j := range input.numbers2 {
			if input.numbers1[i] == input.numbers2[j] {
				appearances++
			}
		}
		score += input.numbers1[i] * appearances
	}

	return score
}

func calculateDistancesSum(distances []int) int {
	sum := 0
	for i := range distances {
		sum += distances[i]
	}

	return sum
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
		a, b := numbersFromLine(line)
		input.numbers1 = append(input.numbers1, a)
		input.numbers2 = append(input.numbers2, b)
	}

	return input
}

func numbersFromLine(line string) (int, int) {
	x := 0
	i := 0

	for ; i < len(line) && line[i] != ' '; i++ {
		digit := line[i] - 48
		x = x*10 + int(digit)
	}
	a := x

	for ; i < len(line) && line[i] == ' '; i++ {
	}

	x = 0
	for ; i < len(line) && line[i] != ' '; i++ {
		digit := line[i] - 48
		x = x*10 + int(digit)
	}
	b := x

	return a, b
}
