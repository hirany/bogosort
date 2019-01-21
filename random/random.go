package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	var len, maxNum int
	var fileName string

	fmt.Print("file name <= ")
	_, err := fmt.Scanf("%s\n", &fileName)
	checkError(err)
	fmt.Print("len <= ")
	_, err = fmt.Scanf("%d\n", &len)
	checkError(err)
	fmt.Print("num <= ")
	_, err = fmt.Scanf("%d\n", &maxNum)
	checkError(err)

	list := makeList(len, maxNum)
	writeIntFile(fileName, list)

}

func makeList(len, maxNum int) (list []int) {

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		list = append(list, rand.Intn(maxNum))
	}
	return

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
