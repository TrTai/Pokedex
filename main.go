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
			inp := cleanInput(scan.Text())
			if cmd, ok := cmdMap[inp[0]]; ok {
				err := cmd.callback()
				if err != nil {
					fmt.Println("Error executing command:", err)
				}
			} else {
				fmt.Println("Unknown command:", inp[0])
			}

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
