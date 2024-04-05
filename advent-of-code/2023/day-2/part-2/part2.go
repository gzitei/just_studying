package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func checkErr(err error) {
    if (err != nil) {
        panic(err);
    }
}

func main() {
    var sum, power int;
    file, err := os.Open("../input.txt");
    checkErr(err);
    defer file.Close();
    scanner := bufio.NewScanner(file);

    for scanner.Scan() {
        
        power = 1;

        line := scanner.Text();

        plays := strings.TrimSpace(line[strings.Index(line, ":") + 1:]);
        
        colors := map[string]int{"red":1, "blue":1, "green":1};

        parts := regexp.MustCompile(" *[,;] *").Split(plays, -1);

        for i := 0; i < len(parts); i++ {
            var color string;
            var count int;
            fmt.Sscanf(parts[i], "%d %s", &count, & color);
            if (count > colors[color]) {
                colors[color] = count;
            }
        }

        for _, v := range colors {
            power *= v;
        }

        fmt.Println(plays, power, colors, sum);

        sum += power;
    }

    fmt.Println("Resposta: ", sum)
}
