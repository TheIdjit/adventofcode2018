package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Gonna start")
	fileval := getFile()
	fmt.Println("Got the goods")
	findMatch(fileval)
}

func getFile() []int {
	val := []int{}
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
		val = append(val, l)
	}
	return val
}

func findMatch(fileval []int) {
	freq := 0
	freqList := []int{}
	i := 0
	for {
		if i == 0 {
			freq = freq + fileval[i]
			freqList = append(freqList, freq)
		}
		if i == len(fileval)-1 {
			fmt.Println("EOF")
			i = 0
		}
		fmt.Println(i)
		if i != (len(fileval)-1) || i != 0 {
			freq = freq + fileval[i]
			for _, freqlistval := range freqList {
				if freq == freqlistval {
					fmt.Println(freq)
					return
				}
			}
			freqList = append(freqList, freq)
		}
		i++
	}
	log.Println(freq)
}
