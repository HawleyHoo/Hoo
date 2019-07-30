/*
@Time : 2019-07-25 10:07
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"fmt"
	"github.com/mgutz/ansi"
)

func example() {
	// colorize a string, SLOW
	msg := ansi.Color("foo", "red+b:white")
	fmt.Println(msg)
	// create a closure to avoid recalculating ANSI code compilation
	phosphorize := ansi.ColorFunc("green+h:black")
	msg = phosphorize("Bring back the 80s!")
	msg2 := phosphorize("Look, I'm a CRT!")

	// cache escape codes and build strings manually
	lime := ansi.ColorCode("green+h:black")
	reset := ansi.ColorCode("reset")

	fmt.Println(lime, "Bring back the 80s!", reset, msg2)
}
func main() {
	example()

	/*
		Colors

			black
			red
			green
			yellow
			blue
			magenta
			cyan
			white

		Attributes

			b = bold foreground
			B = Blink foreground
			u = underline foreground
			h = high intensity (bright) foreground, background
			i = inverse
	*/
	s := "200"
	fmt.Println("-------")
	fmt.Println(ansi.Color(s, "green"))
	fmt.Println(ansi.Color(s, "black"))
	fmt.Println(ansi.Color(s, "yellow"))
	fmt.Println(ansi.Color(s, "blue"))
	fmt.Println(ansi.Color(s, "magenta"))
	fmt.Println(ansi.Color(s, "cyan"))
	fmt.Println(ansi.Color(s, "white"))

	fmt.Println("-------")
	fmt.Println(ansi.Color("200", "off"))
	fmt.Println(ansi.Color(s, "red"))           // red
	fmt.Println(ansi.Color(s, "red+b"))         // red bold
	fmt.Println(ansi.Color(s, "red+B"))         // red blinking
	fmt.Println(ansi.Color(s, "red+u"))         // red underline
	fmt.Println(ansi.Color(s, "red+bh"))        // red bold bright
	fmt.Println(ansi.Color(s, "red:white"))     // red on white
	fmt.Println(ansi.Color(s, "red+b:white+h")) // red bold on white bright
	fmt.Println(ansi.Color(s, "red+B:white+h")) // red blink on white bright

	fmt.Println(ansi.Color(s, "red+b:white+h:0")) // red blink on white bright
	fmt.Println(ansi.Color(s, "red:white"))       // red blink on white bright
	fmt.Println(ansi.Color(s, "green:white"))     // red blink on white bright

}
