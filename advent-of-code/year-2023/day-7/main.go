package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
)

type Hand struct {
	cards      string
	power      int
	tiebreaker int
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
		return h1.tiebreaker >= h2.tiebreaker
	} else {
		return false
	}
}

func newHand(str string) Hand {
	var cards string
	var bid int
	fmt.Sscanf(str, "%s %d", &cards, &bid)
	power, tiebr := handPower(cards)
	return Hand{
		cards:      cards,
		bid:        bid,
		power:      power,
		tiebreaker: tiebr,
	}
}

func newHandV2(str string) Hand {
	var cards string
	var bid int
	fmt.Sscanf(str, "%s %d", &cards, &bid)
	power, tiebr := handPowerV2(cards)
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

func handPowerV2(str string) (int, int) {
	var power, tiebreaker int
	var greater, current Card
	jokerIdx := strings.Index(str, "J")
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
		card := len(cardPower) - (slices.Index(cardPower, s) + 1)
		slcPower[idx] += 1
		slcTiebr[i] = card
		current = Card{
			index: idx,
			value: s,
			power: card,
		}
		if i == 0 {
			greater = current
		} else if current.index != greater.index {
			if slcPower[current.index] > slcPower[greater.index] || (slcPower[current.index] == slcPower[greater.index] && current.power > greater.power) {
				greater = current
			}
		}
	}
	if jokerIdx >= 0 {
		if greater.index != jokerIdx {
			slcPower[greater.index] += slcPower[jokerIdx]
			slcPower[jokerIdx] = 0
		}
	}
	slices.Sort(slcPower)
	slices.Reverse(slcTiebr)
	for i := range sz {
		n := n * i
		power += slcPower[i] * int(math.Pow10(n))
		tiebreaker += slcTiebr[i] * int(math.Pow10(n))
	}
	return power, tiebreaker
}

func handPower(str string) (int, int) {
	var power, tiebreaker int
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
		card := len(cardPower) - (slices.Index(cardPower, s) + 1)
		slcPower[idx] += 1
		slcTiebr[i] = card
	}
	slices.Sort(slcPower)
	slices.Reverse(slcTiebr)
	for i := range sz {
		n := n * i
		power += slcPower[i] * int(math.Pow10(n))
		tiebreaker += slcTiebr[i] * int(math.Pow10(n))
	}
	return power, tiebreaker
}

var wFile *os.File

var cardPower []string = strings.Split("A,K,Q,J,T,9,8,7,6,5,4,3,2", ",")

var n int = int(math.Log10(float64(len(cardPower)))) + 1

func main() {
	var err error
	args := os.Args
	fileName, part := "", ""
	wFile, err = os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer wFile.Close()
	fmt.Sscanf(args[1], "%s %s", &fileName, &part)
	if strings.Contains(part, ".txt") {
		part, fileName = fileName, part
	}
	content := readFile(fileName)
	var choice int
	re := regexp.MustCompile(`[0-9]`)
	part = re.FindString(part)
	fmt.Sscanf(part, "%d", &choice)
	switch choice {
	case 1:
		{
			part1(content)
		}
	case 2:
		{
			part2(content)
		}
	}
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
	// fmt.Println(*i, t.data.cards, t.data.power, t.data.tiebreaker)
	_, err := fmt.Fprintf(wFile, "%d, %s, %d, %d, %d\n", *i, t.data.cards, t.data.bid, t.data.power, t.data.tiebreaker)
	if err != nil {
		fmt.Println(err)
	}
	multiply(t.right, i, s)
}

func part1(content []byte) {
	var rank, sum int
	hands := strings.Split(strings.Trim(string(content), "\n"), "\n")
	hand := newHand(hands[0])
	tr := createNode(hand)
	for i := 1; i < len(hands); i++ {
		thisHand := newHand(strings.TrimSpace(hands[i]))
		tr = tr.add(thisHand)
	}
	multiply(tr, &rank, &sum)
	fmt.Println("Resulting sum is:", sum)
}

func part2(content []byte) {
	var rank, sum int
	j := slices.Index(cardPower, "J")
	for i := j; i < len(cardPower)-1; i++ {
		cardPower[i] = cardPower[i+1]
	}
	cardPower[len(cardPower)-1] = "J"
	hands := strings.Split(strings.Trim(string(content), "\n"), "\n")
	hand := newHandV2(hands[0])
	fmt.Println(cardPower)
	tr := createNode(hand)
	for i := 1; i < len(hands); i++ {
		thisHand := newHandV2(strings.TrimSpace(hands[i]))
		tr = tr.add(thisHand)
	}
	multiply(tr, &rank, &sum)
	fmt.Println("Resulting sum is:", sum)
}
