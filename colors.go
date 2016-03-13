package main

import (
	"github.com/fatih/color"
)

var Fbb = color.New(color.BgRed, color.FgYellow).SprintFunc()
var Red = color.New(color.BgBlack, color.FgRed).SprintFunc()
var Green = color.New(color.BgBlack, color.FgGreen).SprintFunc()
var BlueU = color.New(color.BgBlack, color.FgHiBlue, color.Bold, color.Underline).SprintFunc()
var ItemC = color.New(color.BgBlack, color.FgHiWhite, color.Bold).SprintFunc()
var attrC = color.New(color.BgBlack, color.FgMagenta, color.Underline).SprintFunc()
var Spc = color.New(color.BgBlack, color.FgYellow).SprintFunc()
