package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hfgo/average2"
	"github.com/hfgo/commandline"
)

func main() {
	fmt.Println("")
	fmt.Println(os.Args[1:])
	inputSlice := os.Args[1:]
	floatList, err := commandline.GetFloatCmd(inputSlice)
	if err != nil {
		log.Fatal(err)
	} else {
		average, err := average2.GetAverageFloat(floatList...)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\nAverage is %.2f", average)
		}

	}

}
