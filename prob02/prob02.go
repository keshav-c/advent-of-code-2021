package main

import (
    "fmt"
    "github.com/keshav-c/advent-of-code-2021/util"
    "os"
)

func main() {
    url := "https://adventofcode.com/2021/day/2/input"
    problemLines := util.GetProblem(url)

    resultA := calcA(problemLines)
    fmt.Fprintf(os.Stdout, "Final horizontal x depth: %d\n", resultA)

    resultB := calcB(problemLines)
    fmt.Fprintf(os.Stdout, "True horizontal x depth: %d\n", resultB)
}

