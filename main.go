package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/andmorefine/golang-template/infrastructure"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
func topHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	sum := Average(10000, rand.Intn(10000))
	body := strconv.Itoa(sum)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Top", body)
}

func main() {
	infrastructure.Handle()
	// p := &Page{Title: "test", Body: []byte("this is test page")}
	// p.save()
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/", topHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
