package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "github.com/keshav-c/advent-of-code-2021/util"
)

type position struct {
    X int
    Y int
    aim int
}

func (p *position) executeInstruction(action string, value int) {
    switch action {
        case "forward":
            p.X += value
            p.Y += value * p.aim
        case "up": p.aim -= value
        case "down": p.aim += value
    }
}

func executeInstructions(instructions []string) position {
    var current position
    for _, instruction := range instructions {
        instructionTokens := strings.Split(instruction, " ")
        action := instructionTokens[0]
        value, err := strconv.Atoi(instructionTokens[1])
        util.Check(err, "Atoi")
        current.executeInstruction(action, value)
    }
    return current
}

func calcB(input []string) int {
    finalPosition := executeInstructions(input)

    fmt.Fprintf(os.Stdout, "Final position: %+v\n", finalPosition)
    return finalPosition.X * finalPosition.Y
}

