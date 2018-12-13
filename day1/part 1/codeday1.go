package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	freq := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	read := bufio.NewReader(file)
	for {
		freqMod, _, err := read.ReadLine()
		if err != nil {
			break
		}
		l, _ := strconv.Atoi(string(freqMod))
		freq = freq + l
	}
	fmt.Print(freq)
}
