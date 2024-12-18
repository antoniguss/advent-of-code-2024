package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "./input.txt"
	f, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(f)

	var reports [][]uint

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		report := make([]uint, len(fields))

		for i, field := range fields {
			value, err := strconv.Atoi(field)
			check(err)
			report[i] = uint(value)
		}
		reports = append(reports, report)

	}

	safe := 0
	for _, report := range reports {
		if processReport(report) {
			safe++
		}
	}

	fmt.Printf("Number of safe: %d\n", safe)

	//--- Part 2 ---

	safe = 0
	for _, report := range reports {
		if processReportWithDampener(report) {
			safe++
		}
	}
	fmt.Printf("Number of safe with dampener used: %d\n", safe)
	err = f.Close()
	check(err)

}

func processReportWithDampener(report []uint) bool {
	if processReport(report) {
		return true
	}
	// Else try removing elements to see if it fixes it
	for i := range len(report) {
		var newReport []uint
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)

		if processReport(newReport) {
			return true
		}
	}
	return false

}

func processReport(report []uint) bool {
	ascending := report[1] > report[0]

	if ascending {
		for i := range len(report) - 1 {
			if report[i] >= report[i+1] {
				return false
			}

			if report[i+1]-report[i] > 3 {
				return false
			}

		}

	} else {

		for i := range len(report) - 1 {
			if report[i] <= report[i+1] {
				return false
			}

			if report[i]-report[i+1] > 3 {
				return false
			}

		}
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
