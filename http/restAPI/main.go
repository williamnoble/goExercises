package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

// re-wrote restAPI-extended to slightly improve coding standard.
// further improvements would include - graceful shutdown, create a &http.Server{}, move handlers to separate
// package or struct within app e.g. app.handlers, move err to var ErrUnmarshalerr..

type Article struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
}

type Articles []Article

type application struct {
	Articles Articles
	Routes   *mux.Router
}

func main() {
	fmt.Println("Rest API v0.1 - Mux Router")
	articles := Articles{
		{ID: "1", Title: "Hello", Description: "Article Description", Content: "Article Content"},
		{ID: "2", Title: "Hello 2", Description: "Article Description", Content: "Article Content"},
	}

	App := application{
		Articles: articles,
	}
	App.Routes = App.routes()
	log.Fatal(http.ListenAndServe(":4000", App.Routes))
}

func (a *application) listArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing all articles")
	articles := a.Articles
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		fmt.Println(err)
	}
}

func (a *application) deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range a.Articles {
		if article.ID == id {
			// 1) append everything upto but not including the index
			//2) append everything after the index.
			a.Articles = append(a.Articles[:index], a.Articles[index+1:]...)
		}

	}
}

func (a *application) getArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, article := range a.Articles {
		if article.ID == id {
			if err := json.NewEncoder(w).Encode(&article); err != nil {
				fmt.Println("error", err)
			}
		}
	}
}

func (a *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	bytes, _ := io.ReadAll(r.Body)
	var article Article
	if err := json.Unmarshal(bytes, &article); err != nil {
		fmt.Println("error whilst unmarshalling the response body")
	}

	a.Articles = append(a.Articles, article)
	if err := json.NewEncoder(w).Encode(article); err != nil {
		fmt.Println("error whilst attempting to marshal article body")
	}
}

func (a *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "/Index Page")
}

func (a *application) routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", a.indexHandler)
	router.HandleFunc("/all", a.listArticlesHandler)
	router.HandleFunc("/article", a.createArticleHandler).Methods(http.MethodPost)
	router.HandleFunc("/article/{id}", a.deleteArticleHandler).Methods(http.MethodDelete)
	router.HandleFunc("/article/{id}", a.getArticleHandler).Methods(http.MethodGet)
	//router.HandleFunc("/article/{id}", createArticleHAndler).Methods(http.MethodPost)
	return router
}
