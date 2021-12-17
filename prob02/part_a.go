package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "github.com/keshav-c/advent-of-code-2021/util"
)

type move struct {
    X int
    Y int
}

func parseInstruction(instruction string) move {
    instructionTokens := strings.Split(instruction, " ")
    direction := instructionTokens[0]
    distance, err := strconv.Atoi(instructionTokens[1])
    util.Check(err, "Atoi")
    var parsed move
    switch direction {
        case "forward": parsed.X = distance
        case "up": parsed.Y = -distance
        case "down": parsed.Y = distance
    }
    return parsed
}

func transform(rawLines []string) []move {
    instructions := make([]move, 0)
    for _, line := range rawLines {
        instructions = append(instructions, parseInstruction(line))
    }
    return instructions
}


func (p *move) Add(v move) {
    p.X += v.X
    p.Y += v.Y
}

func calcA(input []string) int {
    instructions := transform(input)
    var position move
    for _, instruction := range instructions {
        position.Add(instruction)
    }

    fmt.Fprintf(os.Stdout, "Final position: %+v\n", position)
    return position.X * position.Y
}
