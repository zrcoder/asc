package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

var fonts = map[string]bool{
	"big":         true,
	"block":       true,
	"chunky":      true,
	"coinstak":    true,
	"colossal":    true,
	"cricket":     true,
	"cyberlarge":  true,
	"cybermedium": true,
	"doh":         true,
	"doom":        true,
	"isometric1":  true,
	"isometric3":  true,
	"larry3d":     true,
	"marquee":     true,
	"ogre":        true,
	"pawp":        true,
	"puffy":       true,
	"rectangles":  true,
	"rounded":     true,
	"slant":       true,
	"small":       true,
	"standard":    true,
	"starwars":    true,
	"stop":        true,
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		showHelp()
		return
	}
	font := ""
	if args[0] == "-f" && len(args) > 1 {
		font = args[1]
		args = args[2:]
	}
	if len(args) == 0 || font != "" && !fonts[font] {
		showHelp()
		return
	}
	figure.NewFigure(strings.Join(args, " "), font, true).Print()
}

func showHelp() {
	figure.NewFigure("asc", "", true).Print()

	info := `
 print text as ascii art
 
 USAGE:
	asc [-f xxx] [text]
 OPTIONS:
	-f, font(default is "")

 SUPPORTED FONTS:
	big
	block
	chunky
	coinstak
	colossal
	cricket
	cyberlarge
	cybermedium
	doh
	doom
	isometric1
	isometric3
	larry3d
	marquee
	ogre
	pawp
	puffy
	rectangles
	rounded
	slant
	small
	standard
	starwars
	stop
	`
	fmt.Println(info)
}
