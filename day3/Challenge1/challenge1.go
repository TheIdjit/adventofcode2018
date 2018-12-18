package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type vibre struct {
	ID        int
	LocationX int
	LocationY int
	sizeX     int
	sizeY     int
	endX      int
	endY      int
}

func main() {
	start := time.Now()
	defer func() {
		log.Printf("It took: %v", time.Since(start))
	}()
	vibres := []*vibre{}
	fileval := getFile()
	for _, s := range fileval {
		v := newVibre(s)
		vibres = append(vibres, v)
	}
	x := makeMatrix(vibres)
	overlap(vibres, x)
}

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
	return val
}

func newVibre(str string) *vibre {
	v := &vibre{}
	data := (strings.Split(str, " "))
	for j := 0; j <= 3; j++ {
		switch j {
		case 0:
			val, _ := strconv.Atoi(strings.Trim(data[j], "#"))
			v.ID = val
		case 2:
			val := strings.Trim(data[j], ":")
			test := strings.Split(val, ",")
			valX, err := strconv.Atoi(test[0])
			if err != nil {
				log.Println(err)
			}
			valY, err := strconv.Atoi(test[1])
			if err != nil {
				log.Println(err)
			}
			v.LocationX = valX
			v.LocationY = valY
		case 3:
			val := strings.Split(data[j], "x")
			valX, err := strconv.Atoi(val[0])
			if err != nil {
				log.Println(err)
			}
			valY, err := strconv.Atoi(val[1])
			if err != nil {
				log.Println(err)
			}
			v.sizeX = valX
			v.sizeY = valY
		}
	}
	v.endX = v.sizeX + v.LocationX
	v.endY = v.sizeY + v.LocationY
	return v
}

func makeMatrix(v []*vibre) [][]int {
	maxX, maxY := getMax(v)
	x := make([][]int, maxX)
	for i, _ := range x {
		y := make([]int, maxY)
		x[i] = y
	}
	return x
}

func getMax(v []*vibre) (maxX, maxY int) {
	for _, val := range v {
		if maxX < val.endX {
			maxX = val.endX
		}
		if maxY < val.endY {
			maxY = val.endY
		}
	}

	return
}

func overlap(v []*vibre, x [][]int) {
	count := 0
	for _, val := range v {
		for i := val.LocationX; i < val.endX; i++ {
			for j := val.LocationY; j < val.endY; j++ {
				x[i][j] = x[i][j] + 1
				if x[i][j] == 2 {
					count++

				}
			}
		}
	}
	fmt.Println(count)
}
