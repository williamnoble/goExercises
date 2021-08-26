package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Posts = []Post{
	Post{Id: 1, Content: "Hello World", Author: "William"},
	Post{Id: 2, Content: "Goodbye Cruel World", Author: "Gabriel"},
	Post{Id: 3, Content: "Hey there Delila", Author: "Old Tom"},
	Post{Id: 4, Content: "Luke, I am your father", Author: "William"},
}

func main() {

	// csv writer

	f, _ := os.Create("f.csv")
	defer f.Close()

	writer := csv.NewWriter(f)

	for _, post := range Posts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			log.Fatal("encountered an error when trying to write to csv")
		}
	}

	// Write buffered data to the underlying writer
	writer.Flush()

	f, _ = os.Open("f.csv")
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	records, _ := reader.ReadAll()

	var posts []Post

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		post := Post{
			Id:      int(id),
			Content: record[1],
			Author:  record[2],
		}
		posts = append(posts, post)
	}

	fmt.Println(posts)

}
