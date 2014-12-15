package class

import (
	"github.com/metakeule/goh4"
)

func color(name string) goh4.Class {
	return class("color-" + name)
}

func background(name string) goh4.Class {
	return class("background-" + name)
}

var (
	ColorBlue   = color("blue")
	ColorGreen  = color("green")
	ColorRed    = color("red")
	ColorYellow = color("yellow")
	ColorOrange = color("orange")
	ColorPink   = color("pink")
	ColorPurple = color("purple")
	ColorGrey   = color("grey")
	ColorBlack  = color("black")
	ColorWhite  = color("white")

	BackgroundBlue   = background("blue")
	BackgroundGreen  = background("green")
	BackgroundRed    = background("red")
	BackgroundYellow = background("yellow")
	BackgroundOrange = background("orange")
	BackgroundPink   = background("pink")
	BackgroundPurple = background("purple")
	BackgroundGrey   = background("grey")
	BackgroundBlack  = background("black")
	BackgroundWhite  = background("white")
)
