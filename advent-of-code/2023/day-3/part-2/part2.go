package main

import (
    "bufio"
    "fmt"
    "os"
)

type Number struct {
    value   int
    row     int
    start   int
    end     int
}

type Multiplier struct {
    row         int
    position    int
    left        int
    right       int
}

func isNumber(c byte) bool {
    return int(c) >= 48 && int(c) < 58
}

func hasNext[T string | []string](n int, piece T) bool {
    return n + 1 <= len(piece)
}

func hasPrevious(n int) bool {
    return n - 1 >= 0
}

func touches (n Number, m Multiplier) bool {
    surroundingLines := n.row >= m.row - 1 && n.row <= m.row + 1
    surroundingColumns := m.position >= n.start - 1 && m.position <= n.end
    return surroundingLines && surroundingColumns
}

func main() {
    args := os.Args
    fileName := args[1]
    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    pieces := make([]string, 0)
    currentRow := 0
    sum := 0
    multipliers := make([]Multiplier, 0)
    numbers := make([]Number, 0)
    start := -1
    numStr := ""
    for scanner.Scan() {
        chunk := scanner.Text()
        pieces = append(pieces, chunk)
        for i, v := range chunk {
            if (!isNumber(byte(v)) || !hasNext(i, chunk)) {
                if (start >= 0) {
                    var value int
                    fmt.Sscanf(numStr, "%d", &value)
                    numbers = append(numbers, Number{
                        value: value,
                        row: currentRow,
                        start: start,
                        end: start + len(numStr),
                    })
                    numStr = ""
                    start = -1
                }
                if v == '*' {
                    next := i
                    prev := i
                    if (hasPrevious(i)) {
                        prev = i - 1
                    }
                    if (hasNext(i, chunk)) {
                        next = next + 1
                    }
                    multipliers = append(multipliers, Multiplier{
                        row: currentRow,
                        position: i,
                        left: prev,
                        right: next,
                    })
                }
            } else if (isNumber(byte(v))){
                numStr = numStr + string(v)
                if (start < 0) {
                    start = i
                }
            }
        }
        currentRow++
    }
    for _, m := range multipliers {
        all := make([]Number, 0)
        for _, n := range numbers {
            itTouches := touches(n, m)
            if (itTouches) {
                all = append(all, n)
            }
        }
        for i := 0; i < len(all) - 1; i++ {
            for j := i + 1; j < len(all); j++ {
                mult := (all[i].value * all[j].value)
                fmt.Println(m.row,"x",m.position, "=>" ,all[i].value, "*", all[j].value, "=>", sum, "+", mult)
                sum += mult
            }
        }
    }
    fmt.Println(sum)
}
