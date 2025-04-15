package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scan.Scan() {
			val := cleanInput(scan.Text())
			fmt.Println("Your command was:", val[0])
		}
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	lower := strings.ToLower(text)
	split := strings.Fields(lower)
	if len(split) == 0 {
		return []string{}
	}
	return split
}
