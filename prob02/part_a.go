package main

import (
    "fmt"
    "os"
)

func (p *delta) Add(v delta) {
    p.deltaX += v.deltaX
    p.deltaY += v.deltaY
}

func calcA(input []delta) int {
    var position delta
    for _, instruction := range input {
        position.Add(instruction)
    }

    fmt.Fprintf(os.Stdout, "Final position: %+v\n", position)
    return position.deltaX * position.deltaY
}
