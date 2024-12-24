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
	inputFile = "test1.txt"
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

	for k, v := range keypadSeq {
		fmt.Printf("(%s -> %s): %+v\n", string(k.from), string(k.to), v)
	}
	dirpad := [][]rune{
		{'_', '^', 'A'},
		{'<', 'v', '>'},
	}

	dirpadSeq := getSequences(dirpad)
	for k, v := range dirpadSeq {
		fmt.Printf("(%s -> %s): %+v\n", string(k.from), string(k.to), v)
	}

	for _, code := range codes {
		moveSeq1 := "A"
		for i := 0; i < len(code)-1; i++ {
			from := rune(code[i])
			to := rune(code[i+1])
			moves := keypadSeq[pair{from: from, to: to}]
			moveSeq1 += moves[0]
		}
		moveSeq2 := "A"
		for i := 0; i < len(moveSeq1)-1; i++ {
			from := rune(moveSeq1[i])
			to := rune(moveSeq1[i+1])
			moves := dirpadSeq[pair{from: from, to: to}]
			moveSeq2 += moves[0]
		}
		moveSeq3 := "A"
		for i := 0; i < len(moveSeq2)-1; i++ {
			from := rune(moveSeq2[i])
			to := rune(moveSeq2[i+1])
			moves := dirpadSeq[pair{from: from, to: to}]
			moveSeq3 += moves[0]
		}
		fmt.Println(moveSeq3[1:])
		fmt.Println(moveSeq2[1:])
		fmt.Println(moveSeq1[1:])
		fmt.Println(code[1:])
		codeNum, _ := strconv.Atoi(code[1 : len(code)-1])
		fmt.Println(len(moveSeq3)-1, codeNum)
		result += (len(moveSeq3) - 1) * codeNum
	}

	return result
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
