package main

// calculate the number of windows whose sum is larger than the previous
func calcB(input []int) int {
    result := calcTotalIncreasingWindows(calcWindowSums(input, 3))
    return <-result
}

// Window sum producer
func calcWindowSums(input []int, wSize int) <-chan int {
    out := make(chan int)
    wSum := 0
    go func() {
        for i := 0; i < wSize; i++ {
            wSum += input[i]
        }
        out <- wSum
        for i := wSize; i < len(input); i++ {
            wSum += input[i] - input[i-wSize]
            out <- wSum
        }
        close(out)
    }()
    return out
}

// Window comparing consumer
func calcTotalIncreasingWindows(wSums <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        result := 0
        prev := <-wSums
        for next := range wSums{
            if next > prev {
                result++
            }
            prev = next
        }
        out <- result
        close(out)
    }()
    return out
}

