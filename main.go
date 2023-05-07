package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	fmt.Println("welcome to RETRO-LED-CLOCK")
	type elements [5]string

	//creating the clock digits from zero to nine
	zero := elements{"|||", "| |", "| |", "| |", "|||"}
	one := elements{"|| ", " | ", " | ", " | ", "|||"}
	two := elements{"|||", "  |", "|||", "|  ", "|||"}
	three := elements{"|||", "  |", "|||", "  |", "|||"}
	four := elements{"| |", "| |", "|||", "  |", "  |"}
	five := elements{"|||", "|  ", "|||", "  |", "|||"}
	six := elements{"|||", "|", "|||", "| |", "|||"}
	seven := elements{"|||", "  |", "  |", "  |", "  |"}
	eight := elements{"|||", "| |", "|||", "| |", "|||"}
	nine := elements{"|||", "| |", "|||", "  |", "|||"}

	// the digits array
	digits := [10][5]string{zero, one, two, three, four, five, six, seven, eight, nine}

	screen.Clear() // clear the terminal

	// infinite loop
	for {
		screen.MoveTopLeft() // move the cursor to the top left
		now := time.Now()
		hr, min, sec := now.Hour(), now.Minute(), now.Second()

		// the timer demarcation
		demarcation := elements{"   ", " | ", "   ", " | ", "   "}
		clk := [8][5]string{digits[hr/10], digits[hr%10], demarcation, digits[min/10], digits[min%10], demarcation, digits[sec/10], digits[sec%10]}
		/** since we have from the line first, then we loop through the elements and print it line by line*/
		for line := 0; line < 5; line++ {
			for i := 0; i < len(clk); i++ {
				next := clk[i][line]
				if (i == 2 || i == 5) && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "\t")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}

}
