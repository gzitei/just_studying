package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check_err (err error) {
    if (err != nil) {
        panic(err);
    }
}

func main() {

    sum := 0;

    colors := map[string]int{
        "red": 12,
        "green": 13,
        "blue": 14,
    }

    var lower, higher int;

    for _, v := range colors {
        if (lower == 0 && higher == 0) {
            lower = v;
            higher = v;
        }
        if (v < lower) {
            lower = v;
        }
        if (v > higher) {
            higher = v;
        }
    }

    file, err := os.Open("../input.txt");
    
    check_err(err);

    defer file.Close();

    scanner := bufio.NewScanner(file);
    
    possible := true;

    for scanner.Scan() {

        var game_id int;

        line := scanner.Text();

        fmt.Sscanf(line, "Game %d:", &game_id);

        dot_idx := strings.Index(line, ":");

        plays := strings.TrimSpace(line[dot_idx+1:]);

        arr_plays := regexp.MustCompile(` *[,;] *`).Split(plays, -1);
        

        for i := 0; i < len(arr_plays); i++ {
            var count int;
            var color string;
            fmt.Sscanf(arr_plays[i], "%d %s", &count, &color);
            switch {
                case count > higher: {
                    possible = false;
                }
                case count <= lower: {
                    possible = true;
                }
                default: {
                    possible = (count <= colors[color]);
                }
            }

            if !possible {
                break;
            }
        
        }
    
        fmt.Println(possible, sum, game_id, plays);

        if (possible) {
            sum = sum + game_id;
        }

    }

    fmt.Println(sum);
}
