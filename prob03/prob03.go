package main

import (
    "fmt"
    "github.com/keshav-c/advent-of-code-2021/util"
    "math"
    "os"
    "strconv"
    "strings"
)

func powInt(base int, exp int) int {
    return int(math.Pow(float64(base), float64(exp)))
}

func absInt(val int) int {
    return int(math.Abs(float64(val)))
}

func strToBits(input string) []int {
    bits := make([]int, 0)
    for _, strBit := range strings.Split(input, "") {
        bit, err := strconv.Atoi(strBit)
        util.Check(err, "Atoi")
        bits = append(bits, bit)
    }
    return bits
}

func main() {
    url := "https://adventofcode.com/2021/day/3/input"
    rawInput := util.GetProblem(url)
    // rawInput := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

    resultA := calcA(rawInput)
    fmt.Fprintf(os.Stdout, "Part a answer: %d\n", resultA)

    resultB := calcB(rawInput)
    fmt.Fprintf(os.Stdout, "Part b answer: %d\n", resultB)
}

