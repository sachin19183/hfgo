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
		// now we count the occurrence of each unique expression in the data
		//Method 1 via array slice
		/*
			var names []string
			var countName []int
			for _, line := range strList {
				matched := false
				for i, name := range names {
					if name == line {
						countName[i]++
						matched = true
					}
				}
				if matched == false {
					names = append(names, line)
					countName = append(countName, 1)
				}
			}
		*/

		//method 2 via map

		var candidateCount = make(map[string]int)
		for _, value := range strList {
			candidateCount[value]++
		}
		/*
			for i, value := range names {
				fmt.Println(value, countName[i])
			}
		*/
		for key, value := range candidateCount {
			fmt.Println(key, ":", value)
		}

	}

}
