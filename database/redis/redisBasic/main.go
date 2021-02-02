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
	connURI      = "localhost:32768"
	connProtocol = "tcp"
)

// type Album struct {
// 	Title  string  `redis:"title"`
// 	Artist string  `redis:"artist"`
// 	Price  float64 `redis:"price"`
// 	Likes  int     `redis:"likes"`
// }

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
	http.ListenAndServe(":4000", mux)

}

func handle(w http.ResponseWriter, r *http.Request) {
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

// _, err = conn.Do("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Println("Eletric Ladyland added successfully!!")

// title, err := redis.String(conn.Do("HGET", "album:1", "title"))
// if err != nil {
// 	log.Fatal(err)
// }

// artist, err := redis.String(conn.Do("HGET", "album:1", "artist"))
// if err != nil {
// 	log.Fatal(err)
// }

// price, err := redis.Float64(conn.Do("HGET", "album:1", "price"))
// if err != nil {
// 	log.Fatal(err)
// }

// likes, err := redis.Int(conn.Do("HGET", "album:1", "likes"))
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)

// // reply, err := redis.StringMap(conn.Do("HGETALL", "album:1"))
// // Fetch all album fields with the HGETALL command. Wrapping this
// // in the redis.Values() function transforms the response into type
// // []interface{}, which is the format we need to pass to
// // redis.ScanStruct() in the next step.
// values, err := redis.Values(conn.Do("HGETALL", "album:1"))
// if err != nil {
// 	log.Fatal(err)
// }

// // Create an instance of an Album struct and use redis.ScanStruct()
// // to automatically unpack the data to the struct fields. This uses
// // the struct tags to determine which data is mapped to which
// // struct fields.
// var album Album
// err = redis.ScanStruct(values, &album)
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Printf("%+v", album)

// // album, err := populateAlbum(reply)
// if err != nil {
// 	log.Fatal(err)
// }

// func populateAlbum(reply map[string]string) (*Album, error) {
// 	var err error
// 	album := new(Album)
// 	album.Title = reply["title"]
// 	album.Artist = reply["artist"]
// 	album.Price, err = strconv.ParseFloat(reply["price"], 64)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Similarly, we need to convert the 'likes' value from a string to
// 	// an integer.
// 	album.Likes, err = strconv.Atoi(reply["likes"])
// 	if err != nil {
// 		return nil, err
// 	}
// 	return album, nil
// }
