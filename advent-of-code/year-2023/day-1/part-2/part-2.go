package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    sum := 0;
    
    numbers := map[string]int{
        "zero": 0,
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    n := make([]string, 0, len(numbers));


    for k := range numbers {
        n = append(n, k);
    }

    fmt.Println(n);

    file, err := os.Open("../input.txt");

    if err != nil {
        fmt.Println("errou!", err)
    }
    defer file.Close();
    scanner := bufio.NewScanner(file);

    for scanner.Scan() {

        line := strings.ToLower(scanner.Text());
        fmt.Println(line);       
        for k, v := range numbers {
            line = strings.ReplaceAll(line, fmt.Sprint(v), k);

        }

        strSize := len(line)
        loop1:
        for i := range strSize - 4 {
            piece := line[i : 5+i]
            for j := range n {
                if (strings.Contains(piece, n[j])) {
                    sum += (10*numbers[n[j]])
                    break loop1
                }
            }
        }
        loop2:
        for i := range strSize - 4 {
            piece := line[strSize-5-i : strSize-i]
            for j := range n {
                if (strings.Contains(piece, n[j])) {
                    sum += numbers[n[j]]
                    break loop2
                }
            }
        }

    }

    fmt.Println("Response: ",sum);

}
