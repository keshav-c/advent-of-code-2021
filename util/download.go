// Fetch
package util

import (
    "bufio"
    "fmt"
    "io"
    "net/http"
    "os"
)

func check(err error, checking string) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s: %v\n", checking, err)
        os.Exit(1)
    }
}

func GetProblem(rawUrl string) []string {
    pRequest, err := http.NewRequest("GET", rawUrl, nil)
    check(err, "NewRequest")
    cookie := http.Cookie{
        Name:"session",
        Value: os.Getenv("ADVENT_OF_CODE_2021"),
    }
    pRequest.AddCookie(&cookie)
    client := http.Client{}
    pResp, err := client.Do(pRequest)
    check(err, "fetch")

    defer pResp.Body.Close()
    return getLines(pResp.Body)
}

func getLines(r io.Reader) []string {
    var lines []string
    pScanner := bufio.NewScanner(r)
    for pScanner.Scan() {
        lines = append(lines, pScanner.Text())
    }
    err := pScanner.Err()
    check(err, "Scan problem lines")

    return lines
}

