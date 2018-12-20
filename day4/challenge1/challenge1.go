package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type guardschedule struct {
	timeAsleep      int
	date            string
	guardID         int
	beginsleep      int
	endsleep        int
	minutessleeping [60]int
}

func main() {
	start := time.Now()
	defer func() {
		log.Printf("It took: %v", time.Since(start))
	}()
	gs := []*guardschedule{}
	a := newSched(gs)
	fmt.Println(a)
}

// func (g *guardschedule) String() string {
// 	val := fmt.Sprintf("Timeasleep:", g.timeAsleep)
// 	return val
// }

func getFile() []string {
	val := []string{}
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
		l := string(freqMod)
		val = append(val, l)
	}
	sort.Strings(val)
	return val
}

func newSched(ns []*guardschedule) int {
	fileval := getFile()
	gds := map[int]*guardschedule{}
	checkcheck := false
	gs := &guardschedule{}
	for _, str := range fileval {
		data := (strings.Split(str, "] "))
		val := strings.Trim(data[0], "[")
		val2 := data[1]
		gs.date = val
		if val2 == "wakes up" {
			wutime := strings.Split(gs.date, " ")
			wutime = strings.Split(wutime[1], ":")
			if checkcheck == true {
				if wutime[0] == "00" {
					gs.endsleep, _ = strconv.Atoi(wutime[1])
					setSleeptime(gs)
					maxsleep := gs.endsleep - gs.beginsleep
					gs.timeAsleep = gs.timeAsleep + maxsleep
					checkcheck = false
				}
			}
		}
		if val2 == "falls asleep" {
			fatime := strings.Split(gs.date, " ")
			fatime = strings.Split(fatime[1], ":")
			if fatime[0] == "00" {
				checkcheck = true
				gs.beginsleep, _ = strconv.Atoi(fatime[1])
			}
		}
		if strings.Contains(val2, "Guard") {
			ns = append(ns, gs)
			checkcheck = false
			gs = &guardschedule{}
			gID := strings.Trim(val2, " begins shift")
			gID = strings.Trim(gID, "Guard #")
			valID, _ := strconv.Atoi(gID)
			if g, ok := gds[valID]; ok {
				gs = g
			} else {
				gds[valID] = &guardschedule{}
				gID = strings.Trim(gID, "Guard #")
				gs = gds[valID]
				gs.guardID, _ = strconv.Atoi(gID)
			}
		}
	}
	_, val := getMaxTimeAsleep(ns)
	thing := gds[val]
	fmt.Println(gds)
	answer := getMinuteAsleep(thing)
	answer = answer * val
	fmt.Println(thing.minutessleeping)
	return answer
}

func getMaxTimeAsleep(gs []*guardschedule) (maxAsleep, maxGuard int) {
	for _, val := range gs {
		if maxAsleep < val.timeAsleep {
			maxAsleep = val.timeAsleep
			maxGuard = val.guardID
		}
	}
	return
}

func setSleeptime(gs *guardschedule) {
	for i, vi := range gs.minutessleeping {
		if gs.beginsleep <= i && i < gs.endsleep {
			gs.minutessleeping[i] = vi + 1
			log.Println(gs.minutessleeping[i])
		}
	}
	log.Println("Hoii", gs.guardID)
}

func getMinuteAsleep(gs *guardschedule) (maxValI int) {
	maxVal := 0
	for i, val := range gs.minutessleeping {
		if maxVal < val {
			maxVal = val
			maxValI = i
		}
	}
	return
}
