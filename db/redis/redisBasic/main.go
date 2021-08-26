package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	connURI      = "localhost:6379"
	connProtocol = "tcp"
)

func main() {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(connProtocol, connURI)
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/album", showAlbum)
	mux.HandleFunc("/", handle)
	mux.HandleFunc("/like", addLike)
	log.Println("listening on port 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Println("Err")
	}

}

func handle(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "connected...")
}

func showAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ShowAlbum1")
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(405), 405)
	}

	id := r.URL.Query().Get("id")
	fmt.Println("ShowAlbum ID ", id)
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	bk, err := findAlbum(id)
	if err == ErrNoAlbum {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprint(os.Stdout, "%s by %s: £%.2f [%d likes] \n", bk.Title, bk.Artist, bk.Price, bk.Likes)
	fmt.Fprintf(w, "%s by %s: £%.2f [%d likes] \n", bk.Title, bk.Artist, bk.Price, bk.Likes)
}

func addLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(405), 405)
		return
	}

	id := r.PostFormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err := IncrementLikes(id)
	if err == ErrNoAlbum {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	http.Redirect(w, r, "/album?id="+id, 303)
}
