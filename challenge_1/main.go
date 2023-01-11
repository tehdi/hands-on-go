package main

import (
	"fmt"
	"github.com/jboursiquot/go-proverbs"
)

func main() {
	fmt.Println(getRandomProverb())
}

func getRandomProverb() string {
	/*
		type Proverb struct {
			Saying string
			Link   string
		}
	*/
	proverb := proverbs.Random()
	return fmt.Sprintf("%s (%s)", proverb.Saying, proverb.Link)
}
