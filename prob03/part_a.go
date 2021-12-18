package main

import (
    "fmt"
    "log"
)

type tally []int
type bitsArr []int

func (t tally) update(bits []int) {
    for i, bit := range bits {
        t[i] += 2*bit - 1
    }
}

func checkVal(v int) {
    if v == 0 {
        log.Fatal("No majority bit found!!")
    }
}

func (t tally) getBits() bitsArr {
    var bits = make([]int, len(t))
    for i, v := range t {
        checkVal(v)
        bits[i] = ((v/absInt(v))+1)/2
    }
    return bits
}

func (b bitsArr) getInverse() bitsArr {
    var inv = make([]int, len(b))
    for i, v := range b {
        inv[i] = 1 - v
    }
    return inv
}

func (b bitsArr) getDecimal() int {
    val := 0
    for i, v := range b {
        val += powInt(2, len(b)-1-i) * v
    }
    return val
}

func calcA(input []string) int {
    var bitTally tally = make([]int, len(input[0]))
    for _, line := range input {
        bits := strToBits(line)
        bitTally.update(bits)
    }
    gammaBits := bitTally.getBits()
    gamma := gammaBits.getDecimal()
    epsilonBits := gammaBits.getInverse()
    epsilon := epsilonBits.getDecimal()
    powerFactor := gamma * epsilon
    fmt.Printf("gamma=%d; epsilon=%d; power_consumption=%d\n", gamma, epsilon, powerFactor)
    return powerFactor
}
