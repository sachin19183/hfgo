package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hfgo/datafile"
)

func main() {
	fmt.Println(len(os.Args))
	if len(os.Args) <= 1 {
		fmt.Println(" Pls provide the filename which you want to read from")
		os.Exit(1)
	}
	fmt.Println("")
	fmt.Println(os.Args[1])
	filename := os.Args[1]
	strList, err := datafile.GetString(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(strList)

	}

}
