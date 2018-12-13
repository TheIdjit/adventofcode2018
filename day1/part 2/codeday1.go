package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	check := false
	readval := []int{}
	val := []int{}
	freq := 0
	n := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	read := bufio.NewReader(file)

	for {
		freqMod, _, err := read.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				n = 0
				fmt.Println("EOF")
				time.Sleep(500 * time.Millisecond)
				n = 0
			}
			continue
		}

		// fmt.Println(freq)
		for i, _ := range val {
			if freq == val[i] {
				check = true
				fmt.Println(freq)
				return
			}
		}

		l, _ := strconv.Atoi(string(freqMod))
		readval = append(readval, l)
		val = append(val, freq)
		if check == true {
			break
		}
		freq = freq + readval[n]
		fmt.Println(freq)
		n++
	}
}
