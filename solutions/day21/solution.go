package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Antyhot/advent-of-code-24/util"
	"github.com/Antyhot/advent-of-code-24/util/queue"
)

const (
	inputFile = "input.txt"
)

func main() {

	fmt.Println("Advent of Code - Day 21") // Placeholder for day number

	file, err := os.Open(inputFile)
	check(err)

	scanner := bufio.NewScanner(file)
	codes := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		codes = append(codes, "A"+line)
	}

	//--- Part 1 ---
	res1 := part1(codes)
	fmt.Printf("Part1: %d\n", res1)

	//--- Part 2 ---

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func part1(codes []string) (result int) {

	keypad := [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{'_', '0', 'A'},
	}

	keypadSeq := getSequences(keypad)

	// for k, v := range keypadSeq {
	// 	fmt.Printf("(%s -> %s): %+v\n", string(k.from), string(k.to), v)
	// }
	dirpad := [][]rune{
		{'_', '^', 'A'},
		{'<', 'v', '>'},
	}

	dirpadSeq := getSequences(dirpad)
	_ = dirpadSeq

	for _, code := range codes {
		// fmt.Printf("Code: %s\n", code[1:])
		robot1 := generatePossibleStrings(code, keypadSeq)

		possible_robot2 := make([]string, 0)
		for _, seq := range robot1 {
			possible_robot2 = append(possible_robot2, generatePossibleStrings("A"+seq, dirpadSeq)...)
		}
		minlen := -1
		var robot2 []string
		for _, seq := range possible_robot2 {
			if len(seq) < minlen || minlen == -1 {
				robot2 = make([]string, 0)
				minlen = len(seq)
			}

			if len(seq) == minlen {
				robot2 = append(robot2, seq)
				continue
			}
		}

		possible_robot3 := make([]string, 0)
		for _, seq := range robot2 {
			possible_robot3 = append(possible_robot3, generatePossibleStrings("A"+seq, dirpadSeq)...)
		}
		minlen = -1
		for _, seq := range possible_robot3 {
			if len(seq) < minlen || minlen == -1 {
				minlen = len(seq)
			}
		}
		codeNum, _ := strconv.Atoi(code[1 : len(code)-1])
		result += minlen * codeNum
	}

	return result
}

func generatePossibleStrings(line string, seq map[pair][]string) []string {
	options := make([][]string, 0)

	for i := 0; i < len(line)-1; i++ {
		from := rune(line[i])
		to := rune(line[i+1])
		moves := seq[pair{from: from, to: to}]
		options = append(options, moves)
	}

	return product(options)
}

func product(list [][]string) []string {
	if len(list) == 0 {
		return []string{""}
	}

	// Initialize the output with an empty string
	output := []string{""}

	// Iterate over each slice in the input list
	for _, opts := range list {
		var temp []string
		// Generate new combinations by appending each option to existing combinations
		for _, prefix := range output {
			for _, opt := range opts {
				temp = append(temp, prefix+opt)
			}
		}
		output = temp
	}

	return output
}

func getSequences(keyPad [][]rune) map[pair][]string {
	positions := make(map[rune]util.Vector)
	for row := range len(keyPad) {
		for col := range len(keyPad[row]) {
			if keyPad[row][col] != '_' {
				positions[keyPad[row][col]] = util.Vector{X: col, Y: row}
			}

		}
	}

	sequences := make(map[pair][]string)
	for x := range positions {
		for y := range positions {
			if x == y {
				sequences[pair{from: x, to: y}] = []string{"A"}
				continue
			}
			possibilities := make([]string, 0)
			queue := queue.NewQueue()
			queue.Enqueue(move{positions[x], ""})
			optimal := 100

			foundOptimal := false
			for !queue.IsEmpty() && !foundOptimal {
				cur := queue.Dequeue().(move)

				options := []move{
					{cur.Above(1), "v"},
					{cur.Right(1), ">"},
					{cur.Below(1), "^"},
					{cur.Left(1), "<"},
				}

				for _, opt := range options {
					if !opt.WithinBounds(len(keyPad[0]), len(keyPad)) || keyPad[opt.Y][opt.X] == '_' {
						continue
					}
					if len(cur.path)+1 > optimal {
						foundOptimal = true
						break
					}
					if keyPad[opt.Y][opt.X] == y {
						optimal = len(cur.path) + 1
						possibilities = append(possibilities, (cur.path + opt.path + "A"))
					} else {
						queue.Enqueue(move{opt.Vector, cur.path + opt.path})
					}
				}

			}

			sequences[pair{from: x, to: y}] = possibilities
		}
	}
	return sequences
}

type move struct {
	util.Vector
	path string
}

type pair struct {
	from rune
	to   rune
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
