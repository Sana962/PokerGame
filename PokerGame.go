package main

import (
	"fmt"
)

var hand [5][3]string
var deck [5][3]string
var play [5][3]string

const ( //Hand  values
	HIGHEST_CARD = 0
	ONE_PAIR
	TWO_PAIRS
	THREE_OF_A_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_A_KIND
	STRAIGHT_FLUSH
)

var hands = []string{"highest-card", "one-pair", "two-pairs", "three-of-a-kind", "straight", "flush", "full-house", "four-of-a-kind", "straight-flush"}

func main() {
	var j int
	fmt.Println("Enter Cards")
	fmt.Scanf("%s %s %s %s %s %s %s %s %s %s", &hand[0], &hand[1], &hand[2], &hand[3], &hand[4], deck[0], deck[1], deck[2], deck[3], deck[4]) //Enter cards
	var best = 0
	for i := 1; i < 32; i++ {
		var k = 0
		for j := 0; j < 5; j++ {
			if i&1<<j != 0 {
				play[j] = hand[j]
			} else {
				play[j] = deck[k]
				k++
			}
		}
		j = calc()

		if j > best {
			best = j
		}
	}

	fmt.Println("Hand:")
	for i := 0; i < 5; i++ {
		fmt.Println("Hand is", hand[i])
	}
	fmt.Println("Deck:")
	for i := 0; i < 5; i++ {
		fmt.Println("Deck is", deck[i])
	}
	fmt.Println("Best Hand Is", hands[best])

}

func flush() bool {
	var c string = play[0][1]

	for i := 1; i < 5; i++ {

		if c != play[i][0] {

			return false
		}
	}
	return true
}
func kind(a int) bool { // 3 of kind have 3 cards of one rank and two cards of two other rank , 2 of kind has to cards of a matching card

	for i := 0; i < 5-a; i++ {
		for j := 1; j < a; j++ {
			if play[i][0] != play[i+j][0] {
				goto label1
			}
		}
		return true
	label1:
	}
	return false
}

func fullhouseKind(a int, b int) bool { //Any three cards of same number ot face value , plus any other two cards of the same number ot face value
	for i := 0; i < 5-a-b; i++ {
		for j := i + a; j <= 5; j++ {
			for k := 1; k < a; k++ {
				if play[i][0] != play[i+k][0] {
					goto label2
				}
			}
			for k := 1; k < b; k++ {
				if play[j][0] != play[j+k][0] {
					goto label2
				}
			}
			return true
		label2:
		}
	}
	return false
}

func calc() int { //Best hand will be selected
	if kind(4) {
		return FOUR_OF_A_KIND
	} else if fullhouseKind(3, 2) || fullhouseKind(2, 3) {
		return FULL_HOUSE
	} else if flush() {
		return FLUSH
	} else if kind(3) {
		return THREE_OF_A_KIND
	} else if fullhouseKind(2, 2) {
		return TWO_PAIRS
	} else if kind(2) {
		return ONE_PAIR
	} else {
		return HIGHEST_CARD
	}
}
