package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"
)

type Hand struct {
	cards      string
	power      int
	tiebreaker []int
	bid        int
}

type Node struct {
	data  Hand
	left  *Node
	right *Node
}

func createNode(data Hand) *Node {
	return &Node{data: data}
}

func (t *Node) add(data Hand) *Node {
	if t == nil {
		return createNode(data)
	} else if data.greaterThan(t.data) {
		t.right = t.right.add(data)
	} else {
		t.left = t.left.add(data)
	}
	return t
}

func showInOrder(t *Node) {
	if t == nil {
		return
	}
	showInOrder(t.left)
	fmt.Print(t.data, " ")
	showInOrder(t.right)
}

func (h1 Hand) greaterThan(h2 Hand) bool {
	if h1.power > h2.power {
		return true
	} else if h1.power == h2.power {
		r := false
		for i := range h1.tiebreaker {
			if h1.tiebreaker[i] == h2.tiebreaker[i] {
				continue
			} else {
				r = h1.tiebreaker[i] > h2.tiebreaker[i]
				break
			}
		}
		return r
	} else {
		return false
	}
}

func newHand(str string, joker string) Hand {
	var cards string
	var bid int
	fmt.Sscanf(str, "%s %d", &cards, &bid)
	power, tiebr := handPower(cards, joker)
	return Hand{
		cards:      cards,
		bid:        bid,
		power:      power,
		tiebreaker: tiebr,
	}
}

type Card struct {
	index int
	value string
	power int
}

func handPower(str, joker string) (int, []int) {
	var power int
	jokerCount := 0
	greaterIdx := -1
	greaterCard := -1
	jokerIdx := -1
	sz := len(str)
	slcPower := make([]int, 0, sz)
	slcTiebr := make([]int, 0, sz)
	for range sz {
		slcPower = append(slcPower, 0)
		slcTiebr = append(slcTiebr, 0)
	}
	for i := 0; i < sz; i++ {
		s := string(str[i])
		idx := strings.Index(str, s)
		card := slices.Index(cardPower, s) + 1
		slcTiebr[i] = card
		if s == joker {
			jokerCount++
			if jokerIdx == -1 {
				jokerIdx = i
			}
			continue
		}
		slcPower[idx] += 1
		if greaterIdx == -1 {
			greaterIdx = idx
			greaterCard = card
		} else if slcPower[idx] >= slcPower[greaterIdx] {
			if slcPower[idx] > slcPower[greaterIdx] || card > greaterCard {
				greaterIdx = idx
				greaterCard = card
			}
		}
	}
	if greaterIdx == -1 {
		slcPower[jokerIdx] = jokerCount
	} else {
		slcPower[greaterIdx] += jokerCount
	}
	slices.Sort(slcPower)
	for i := range sz {
		power += slcPower[i] * int(math.Pow10(i))
	}
	fmt.Println(str, power, slcPower, slcTiebr)
	return power, slcTiebr
}

var cardPower []string = strings.Split("A,K,Q,J,T,9,8,7,6,5,4,3,2", ",")

func main() {
	start := time.Now()
	args := os.Args
	fileName, part, joker := args[1], args[2], args[3]
	content := readFile(fileName)
	var choice int
	re := regexp.MustCompile(`[0-9]`)
	part = re.FindString(part)
	fmt.Sscanf(part, "%d", &choice)
	switch choice {
	case 1:
		{
			part1(content, "")
		}
	case 2:
		{
			part2(content, joker)
		}
	}
	fmt.Println("Ran in:", time.Since(start).Milliseconds())
}

func readFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func multiply(t *Node, i, s *int) {
	if t == nil {
		return
	}
	multiply(t.left, i, s)
	*i += 1
	*s += *i * t.data.bid
	fmt.Println(t.data.cards, "=>", t.data.power, t.data.tiebreaker, "|", t.data.bid, *i)
	multiply(t.right, i, s)
}

func part1(content []byte, joker string) {
	slices.Reverse(cardPower)
	var rank, sum int
	hands := strings.Split(strings.Trim(string(content), "\n"), "\n")
	hand := newHand(hands[0], joker)
	tr := createNode(hand)
	for i := 1; i < len(hands); i++ {
		thisHand := newHand(strings.TrimSpace(hands[i]), joker)
		tr = tr.add(thisHand)
	}
	showInOrder(tr)
	multiply(tr, &rank, &sum)
	fmt.Println("Resulting sum is:", sum)
}

func part2(content []byte, joker string) {
	var rank, sum int
	j := slices.Index(cardPower, joker)
	for i := j; i < len(cardPower)-1; i++ {
		cardPower[i] = cardPower[i+1]
	}
	cardPower[len(cardPower)-1] = joker
	slices.Reverse(cardPower)
	hands := strings.Split(strings.Trim(string(content), "\n"), "\n")
	hand := newHand(hands[0], joker)
	tr := createNode(hand)
	for i := 1; i < len(hands); i++ {
		thisHand := newHand(strings.TrimSpace(hands[i]), joker)
		tr = tr.add(thisHand)
	}
	multiply(tr, &rank, &sum)
	fmt.Println("Resulting sum is:", sum)
}
