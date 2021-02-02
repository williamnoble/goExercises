package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	// ID   string       `json:"ID"`
	Title   string `json:"string"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hit the HomePage")
	fmt.Println("Endpoint hit: Homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Beast Slays Man", Desc: "A Heroic Beast has killed the hunter which hunted him", Content: "Interview with beast to follow!"},
		Article{Title: "Angel finds Love", Desc: "Angel has married Warewolf in extra-ordinary turn of events", Content: "Ex-Rated content cersored!!xxx"},
	}
	fmt.Println("Endpoint hit: returnAllArticles")
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		fmt.Println(err)
	}

}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "%v", key)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10001", myRouter))
}

func main() {
	handleRequests()
}
