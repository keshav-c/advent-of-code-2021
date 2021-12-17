package main

import (
    "fmt"
    "github.com/keshav-c/advent-of-code-2021/util"
    "strconv"
    "strings"
    "os"
)

type delta struct {
    deltaX int
    deltaY int
}

func parseInstruction(instruction string) delta {
    instructionTokens := strings.Split(instruction, " ")
    direction := instructionTokens[0]
    distance, err := strconv.Atoi(instructionTokens[1])
    util.Check(err, "Atoi")
    var parsed delta
    switch direction {
        case "forward": parsed.deltaX = distance
        case "up": parsed.deltaY = -distance
        case "down": parsed.deltaY = distance
    }
    return parsed
}

func transform(rawLines []string) []delta {
    inputData := make([]delta, 0)
    for _, line := range rawLines {
        inputData = append(inputData, parseInstruction(line))
    }
    return inputData
}

func main() {
    url := "https://adventofcode.com/2021/day/2/input"
    problemLines := util.GetProblem(url)
    input := transform(problemLines)

    resultA := calcA(input)
    fmt.Fprintf(os.Stdout, "Final horizontal x depth: %d\n", resultA)

    resultB := calcB()
    fmt.Fprintf(os.Stdout, "Windows increased: %d\n", resultB)
}

