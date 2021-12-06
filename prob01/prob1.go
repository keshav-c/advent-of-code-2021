package main

import (
    "fmt"
    "github.com/keshav-c/advent-of-code-2021/util"
    "os"
)

func main() {
    url := "https://adventofcode.com/2021/day/1/input"
    problemLines := util.GetProblem(url)
    fmt.Fprintf(os.Stdout, "Number of lines: %d\n", len(problemLines))
}

