package main

import (
    "fmt"
    "github.com/keshav-c/advent-of-code-2021/util"
    "strconv"
    "os"
)

func transform(rawLines []string) []int {
    inputData := make([]int, 0)
    for _, line := range rawLines {
        inputVal, err := strconv.Atoi(line)
        util.Check(err, "Atoi")
        inputData = append(inputData, inputVal)
    }
    return inputData
}

func main() {
    url := "https://adventofcode.com/2021/day/1/input"
    problemLines := util.GetProblem(url)
    input := transform(problemLines)

    resultA := calcA(input)
    fmt.Fprintf(os.Stdout, "Numbers increased: %d\n", resultA)

    resultB := calcB(input)
    fmt.Fprintf(os.Stdout, "Windows increased: %d\n", resultB)
}

