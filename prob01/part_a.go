package main

// Calculate number of values in input array larger than the previous value
func calcA(input []int) int {
    result := 0
    for i := 1; i < len(input); i++ {
        if input[i] > input[i-1] {
            result++
        }
    }
    return result
}

