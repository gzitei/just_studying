package main

func minOperations(logs []string) int {
	depth := 0
	for i := 0; i < len(logs); i++ {
		if logs[i] == "../" {
			if depth > 0 {
				depth--
			}
		} else if logs[i] != "./" {
			depth++
		}
	}
	return depth
}

func main() {}

