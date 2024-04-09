package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func checkErr (err error) {
    if (err != nil) {
        panic(err);
    } 
}

func checkArray (str string) int {
    count := 0
    for i := range str {
        curr := str[i]
        n := int(curr) - 48;
        isNumber := (n >=0 && n <=9)
        isDot := curr == byte('.');
        if (!isNumber && !isDot) {
            count++;
        }
    }
    return count;
}

func main () {
    args := os.Args;
    fileName := args[1];
    file, err := os.Open(fileName)
    checkErr(err);
    defer file.Close();
    scanner := bufio.NewScanner(file);
    chunks := make([]string, 0);
    currentRow := 0;
    sum := 0;
    dictNumbers := map[int][](map[string]int){};

    for scanner.Scan() {
        chunk := scanner.Text();
        chunks = append(chunks, chunk);
        numbers := make([]map[string]int, 0, len(chunk));
        currNumber := "";
        start := -1;
        for i := range chunk {
            curr := int(chunk[i]) - 48;
            isLast := i == (len(chunk) -1)
            isNumber := curr >= 0 && curr <= 9
            if (isNumber && start < 0) {
                start = i
            }
            if (isNumber) {
                currNumber += string(chunk[i])
                if (isLast) {
                    var value int;
                    fmt.Sscanf(currNumber, "%d", &value);
                    el := map[string]int{
                        "value": value,
                        "start": start,
                    }
                    numbers = append(numbers, el);
                    start = -1
                    currNumber = ""
                }
            } else if (start >= 0 && (isLast || !isNumber)) {
                var value int;
                fmt.Sscanf(currNumber, "%d", &value);
                el := map[string]int{
                    "value": value,
                    "start": start,
                }
                numbers = append(numbers, el);
                start = -1
                currNumber = ""
            }
        }
        numbers = slices.Clip(numbers);
        dictNumbers[currentRow] = numbers;
        currentRow++;
    }
    for row := range currentRow {
        numbers := dictNumbers[row]
        fmt.Println(row, numbers)
        lines := make([]string, 0, 3);
        if (row - 1 >= 0) {
            lines = append(lines, chunks[row - 1])
        }
        currentLine := chunks[row];
        lines = append(lines, currentLine)
        if (row + 1 < len(chunks)){
            lines = append(lines, chunks[row + 1])
        }
        for i := range numbers {
            value := numbers[i]["value"]
            start :=  numbers[i]["start"]
            str := fmt.Sprint(value);
            end := start + len(str);
            spanStart := start - 1;
            if (spanStart < 0) {
                spanStart = 0
            }
            spanEnd := end + 1;
            if (spanEnd > len(currentLine)) {
                spanEnd = len(currentLine)
            }
            for l := range lines {
                span := lines[l][spanStart:spanEnd]
                fmt.Println(row, lines[l], span, value)
                mult := checkArray(span)
                numbers[i]["mult"] = mult
                sum += (mult * value)
            }
            fmt.Println(numbers[i], spanStart, spanEnd)
        }
    }
    fmt.Println(sum);
}
