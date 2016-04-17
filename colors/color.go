// color. Here find functions that apply color to text.  Thanks to the nifty
// color package by fatih (github.com/fatih/color).  And I'm actually using
// his vim-go package, although with some modifications to write this :)
package color

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

// normal colors
var Red = color.New(color.FgRed).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
var Blue = color.New(color.FgBlue).SprintFunc()
var Magenta = color.New(color.FgMagenta).SprintFunc()
var Cyan = color.New(color.FgCyan).SprintFunc()
var White = color.New(color.FgWhite).SprintFunc()
var Black = color.New(color.FgBlack).SprintFunc()

// with black background
var RedB = color.New(color.BgBlack, color.FgRed).SprintFunc()
var GreenB = color.New(color.BgBlack, color.FgGreen).SprintFunc()
var YellowB = color.New(color.BgBlack, color.FgYellow).SprintFunc()
var BlueB = color.New(color.BgBlack, color.FgBlue).SprintFunc()
var MagentaB = color.New(color.BgBlack, color.FgMagenta).SprintFunc()
var CyanB = color.New(color.BgBlack, color.FgCyan).SprintFunc()
var WhiteB = color.New(color.BgBlack, color.FgWhite).SprintFunc()

// normal with underline
var RedU = color.New(color.FgRed, color.Underline).SprintFunc()
var GreenU = color.New(color.FgGreen, color.Underline).SprintFunc()
var YellowU = color.New(color.FgYellow, color.Underline).SprintFunc()
var BlueU = color.New(color.FgBlue, color.Underline).SprintFunc()
var MagentaU = color.New(color.FgMagenta, color.Underline).SprintFunc()
var CyanU = color.New(color.FgCyan, color.Underline).SprintFunc()
var WhiteU = color.New(color.FgWhite, color.Underline).SprintFunc()

// black background and underline
var RedBU = color.New(color.BgBlack, color.FgRed, color.Underline).SprintFunc()
var GreenBU = color.New(color.BgBlack, color.FgGreen, color.Underline).SprintFunc()
var YellowBU = color.New(color.BgBlack, color.FgYellow, color.Underline).SprintFunc()
var BlueBU = color.New(color.BgBlack, color.FgBlue, color.Underline).SprintFunc()
var MagentaBU = color.New(color.BgBlack, color.FgMagenta, color.Underline).SprintFunc()
var CyanBU = color.New(color.BgBlack, color.FgCyan, color.Underline).SprintFunc()
var WhiteBU = color.New(color.BgBlack, color.FgWhite, color.Underline).SprintFunc()

// Special non specific color forma
var AttrC = color.New(color.BgBlack, color.FgMagenta, color.Underline).SprintFunc()
var Fbb = color.New(color.BgRed, color.FgYellow).SprintFunc()
var Spc = color.New(color.BgBlack, color.FgYellow).SprintFunc()
var FbbBanner = color.New(color.FgRed).SprintFunc()

//var ItemC = color.New(color.BgBlack, color.FgHiWhite, color.Bold).SprintFunc()
var ItemC = color.New(color.FgHiWhite, color.Bold).SprintFunc()

func Clear() {

	fmt.Print("[H[J")
}

// Switcheroo randomizes color in a byte slice
func Switcheroo(s string) string {

	textOut := ""
	rainbow := []func(...interface{}) string{
		Green,
		Blue,
		Red,
		Yellow,
		Cyan,
		Magenta,
		White,
		Black,
	}

	for _, i := range s {
		textOut = textOut + fmt.Sprintf("%s", rainbow[rand.Intn(len(rainbow))](string(i)))
	}
	return textOut
}

// ColorBanner randomizes color of items in byte slice, specifically for logo banner
func ColorBanner(s string) string {
	textOut := ""
	rainbow := []func(...interface{}) string{
		Green,
		Green,
		Green,
		Green,
		Green,
		Green,
		Green,
		Yellow,
		Blue,
	}

	for _, i := range s {
		textOut = textOut + fmt.Sprintf("%s", rainbow[rand.Intn(len(rainbow))](string(i)))
	}
	return textOut
}
