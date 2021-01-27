// Package page deals with reading http get response and calcualting the size of http.body
package page

import (
	"io/ioutil"
	"log"
	"net/http"
)

//Page is a struct which contains a url string and a SIZE parameter of type int
type Paged struct {
	SiteUrl  string
	PageSize int
}

//ResponseSize takes as input a url string and a chanel with type struct Page .It reads the body of the url and calculates the page size in bytes
func ResponseSize(url string, channel chan Paged) {
	var response *http.Response
	//fmt.Println("Getting ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(len(body))
	channel <- Paged{SiteUrl: url, PageSize: len(body)}
}
