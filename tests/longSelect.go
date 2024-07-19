//go:build ignore

package main

import "github.com/connor-ve/survey/v3"

func main() {
	color := ""
	prompt := &survey.Select{
		Message: "Choose a color:",
		Options: []string{
			"a",
			"b",
			"c",
			"d",
			"e",
			"f",
			"g",
			"h",
			"i",
			"j",
		},
	}
	survey.AskOne(prompt, &color)
}
