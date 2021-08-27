package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      string `json:"ID"`
	Title   string `json:"string"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hit the HomePage")
	fmt.Println("Endpoint hit: Homepage")
}

func returnAllArticles(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")
	if err := json.NewEncoder(w).Encode(Articles); err != nil {
		fmt.Println(err)
	}

}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "%v", key)

	for _, article := range Articles {
		if article.ID == key {
			if err := json.NewEncoder(w).Encode(article); err != nil {
				fmt.Println("Error", err)
			}
		}
	}
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of the POST request and return string resp with req body
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	var article Article
	if err := json.Unmarshal(reqBody, &article); err != nil {
		fmt.Println("Error")
	} // pass pointer or duplicates;

	Articles = append(Articles, article)
	if err := json.NewEncoder(w).Encode(article); err != nil {
		fmt.Println("Error")
	}

}

func deleteArticle(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article", CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	fmt.Println("listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
