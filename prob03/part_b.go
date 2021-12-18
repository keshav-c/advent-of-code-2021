package main

import (
    "fmt"
)

type bits []int
type data []bits

func (d data) getFilteringBit(position int, rows []int, criteria string) int {
    var tally, filter int
    for _, rowIndex := range rows {
        tally += 2*d[rowIndex][position] - 1
    }
    switch criteria {
    case "o2":
        if tally >= 0 {
            filter = 1
        } else {
            filter = 0
        }
    case "co2":
        if tally < 0 {
            filter = 1
        } else {
            filter = 0
        }
    }
    return filter
}

func (d data) filterData(position int, filter int, rows []int) []int {
    filteredRows := make([]int, 0)
    for _, rowIndex := range rows {
        if d[rowIndex][position] == filter {
            filteredRows = append(filteredRows, rowIndex)
        }
    }
    return filteredRows
}

func (d data) getRating(criteria string) bits {
    rows := make([]int, len(d))
    for i, _ := range d {
        rows[i] = i
    }
    var position, filter int
    for i:=0; i<100;i++ {
        filter = d.getFilteringBit(position, rows, criteria)
        rows = d.filterData(position, filter, rows)
        fmt.Printf("pos %d filter %d\n", position, filter)
        fmt.Printf("len %d rows %v\n", len(rows), rows)
        if len(rows) == 1 {
            fmt.Printf("Returning...\n")
            return d[rows[0]]
        } else {
            position = (position+1)%len(d[0])
        }
    }
    fmt.Printf("how did i reach here?\n")
    return d[rows[0]]
}

func (b bits) getDecimal() int {
    val := 0
    for i, v := range b {
        val += powInt(2, len(b)-1-i) * v
    }
    return val
}

func calcB(input []string) int {
    var inputData data = make([]bits, 0)
    for _, line := range input {
        bitSequence := strToBits(line)
        inputData = append(inputData, bitSequence)
    }
    o2_rating := inputData.getRating("o2").getDecimal()
    co2_rating := inputData.getRating("co2").getDecimal()
    power_factor := o2_rating*co2_rating
    fmt.Printf("o2_rating=%d; co2_rating=%d; life_support_rating=%d\n",
        o2_rating, co2_rating, power_factor)
    return power_factor
}
