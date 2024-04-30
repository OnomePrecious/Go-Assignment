package main

import "fmt"

func main() {
	var day int
	var giftDay int
	for day = 1; day <= 12; day++ {
		fmt.Printf("On the %s day of christmas, my true love gave to me:\n", getDay(day))

		for giftDay = day; giftDay >= 1; giftDay-- {
			fmt.Println(" " + getGift(giftDay))
		}
		fmt.Println()
	}
}

func getDay(day int) string {
	switch day {
	case 1:
		return "first"
	case 2:
		return "second"
	case 3:
		return "third"
	case 4:
		return "fourth"
	case 5:
		return "fifth"
	case 6:
		return "sixth"
	case 7:
		return "seventh"
	case 8:
		return "eight"
	case 9:
		return "ninth"
	case 10:
		return "tenth"
	case 11:
		return "eleventh"
	case 12:
		return "twelfth"
	default:
		return "Invalid"
	}
}
func getGift(giftDay int) string {
	switch giftDay {
	case 12:
		return "Twelve drummers drumming "
	case 11:
		return "Eleven pipers piping"
	case 10:
		return "Ten lords are leaping"
	case 9:
		return "Nine ladies dancing"
	case 8:
		return "Eight maids are milking"
	case 7:
		return "Seven swarms are swimming"
	case 6:
		return "Six geese are laying"
	case 5:
		return "Five golden rings"
	case 4:
		return "Four calling birds"
	case 3:
		return "Three fresh hens"
	case 2:
		return "Two turtle doves"
	case 1:
		return "A partridge in a pear tree"
	default:
		return "Invalid"

	}
}
