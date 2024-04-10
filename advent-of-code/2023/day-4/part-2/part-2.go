package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "regexp"
    "strings"
)

type Game struct {
    line string
    rpt int
}

func addOrdered(n int, arr []int) []int {
    size := len(arr)
    if (size == 0 || n >= arr[size - 1]) {
        arr = append(arr, n)
    } else {
        for i, v := range arr {
            if (v > n) {
                for j := len(arr) - 1; j >= i; j-- {
                    if (j == len(arr) - 1) {
                        arr = append(arr, arr[j]);
                    } else {
                        arr[j + 1] = arr[j]
                    }
                }
                arr[i] = n
                break
            }
        }
    }
    return arr
}

func isWinningNumber(n int, arr []int) bool {
    var middle int;
    size := float64(len(arr))
    middle = int(math.Floor(size/2.0))
    arrLeft := arr[0:middle]
    arrRight := arr[middle:]
    if size == 0 {
        return false;
    } else if size == 1 {
        return n == arr[0]
    } else if (n >= arrRight[0]) {
        return isWinningNumber(n, arrRight)
    } else {
        return isWinningNumber(n, arrLeft)
    }
}

func main() {
    args := os.Args
    fileName := args[1]
    file,err := os.Open(fileName)
    if (err != nil) {
        fmt.Println(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    points := 0
    sum := 0
    pieces := make([]Game, 0)
    for scanner.Scan() {
        chunk := scanner.Text()
        pieces = append(pieces, Game{
            line: chunk,
            rpt: 1,
        })
    }
    for idx, v := range pieces {
        chunk := v.line
        next := 0
        colonPosition := strings.Index(chunk, ":")
        dividerPosition := strings.Index(chunk, "|")
        winners := chunk[colonPosition+1:dividerPosition-1]
        winners = strings.TrimSpace(winners)
        reSpace := regexp.MustCompile(` {1,}`)
        arrWinners := reSpace.Split(winners, -1)
        game := reSpace.Split(strings.TrimSpace(chunk[dividerPosition+1:]), -1)
        winnerNumbers := make([]int, 0, len(winners))
        for _, v := range arrWinners {
            var value int
            s := strings.TrimSpace(v)
            if (s == ""){
                continue
            }
            fmt.Sscanf(s, "%d", &value)
            winnerNumbers = addOrdered(value, winnerNumbers)
        }
        for _, v := range game {
            var value int
            s := strings.TrimSpace(v)
            if (s == "") {
                continue
            }
            fmt.Sscanf(s, "%d", &value)
            exists := isWinningNumber(value, winnerNumbers)
            if (exists) {
                next++
            }
        }
        for i := 1; i <= next; i ++ {
            pieces[idx + i].rpt += 1*v.rpt
        }
        next = 0
        sum += v.rpt
    }
    for _, p := range pieces {
        fmt.Println(p.rpt, p.line)
        points += p.rpt
    }
    fmt.Println(points, sum);
}
