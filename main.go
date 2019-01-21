package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {

	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("err: args")
		fmt.Println("go run main.go \"input file name\" \"output file name\"")
		os.Exit(1)
	}
	input := flag.Args()[0]
	output := flag.Args()[1]

	list := readIntFile(input)
	count := bogoSort(list)
	fmt.Println("exchange count: ", count*len(list))
	writeIntFile(output, list)

}

func bogoSort(a []int) int {

	count := 0
	for !sort.IntsAreSorted(a) {
		count++
		shuffle(a)
		showCalm(a)
	}
	return count

}

func shuffle(a []int) {

	var k int
	rand.Seed(time.Now().UnixNano())

	for i := range a {
		k = rand.Intn(i + 1)
		a[i], a[k] = a[k], a[i]
	}

}

func showCalm(a []int) {

	for i := range a {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println("")

}

func readIntFile(fileName string) []int {

	var n, tmp int
	var a []int

	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	for i := 0; true; i++ {
		n, _ = fmt.Fscanf(file, "%d\n", &tmp)
		if n == 0 {
			break
		}
		a = append(a, tmp)
	}

	return a

}

func writeIntFile(fileName string, a []int) {

	file, err := os.Create(fileName)
	checkError(err)
	defer file.Close()

	for i := range a {
		fmt.Fprintf(file, "%d\n", a[i])
	}

}

func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}

}
