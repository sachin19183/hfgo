package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type GuestBook struct {
	SignCount int
	Signature []string
}

func getString(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err) // for any other error than "file does not exist"
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	check(scanner.Err())
	return lines
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	//message := []byte("Signature list goes here!!!")
	signatures := getString("signature.txt")
	fmt.Printf("%#v", signatures)
	guestbook := GuestBook{
		SignCount: len(signatures),
		Signature: signatures,
	}
	html, err := template.ParseFiles("view.html") // use the content of this file to create a template
	check(err)
	err = html.Execute(writer, guestbook) //writes the template content to ResponseWriter
	check(err)
}
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("index.html")
	check(err)
	err = html.Execute(writer, nil)
}
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	//_, err := writer.Write([]byte(signature))
	option := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signature.txt", option, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}
func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/index", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	//http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("http/css"))))
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}
