package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

    start := time.Now();
    
    sum := 0;

    file, err := os.Open("input.txt");
    
    if err != nil {
        fmt.Println("errou!", err)
    }

    scanner := bufio.NewScanner(file);

    for scanner.Scan() {

        bytes := scanner.Bytes();

        var pos int;

        for i := 0; i < len(bytes); i ++ {
            curr := int(bytes[i]) - 48;
            if (curr >= 0 && curr < 10) {
                sum += 10*curr;
                pos = i;
                break;
            }
        }

        for i := len(bytes) - 1; i >= pos; i-- {
            curr := int(bytes[i]) - 48;
            if (curr >= 0 && curr < 10) {
                sum += curr;
                break;
            }
        }

    }

    end := time.Now();

    fmt.Println(sum, end.Sub(start));

}
