package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var (
	post1 = Post{Id: 1, Content: "Hello World", Author: "William"}
	post2 = Post{Id: 2, Content: "Goodbye Cruel World", Author: "Gabriel"}
	post3 = Post{Id: 3, Content: "Hey there Delila", Author: "Old Tom"}
	post4 = Post{Id: 4, Content: "Luke, I am your father!", Author: "Dominic"}
)

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	// Create me some post data

	store(post1)
	store(post2)
	store(post3)
	store(post4)
	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["William"] {
		fmt.Println("Author: "+post.Author, " Content: "+post.Content)
	}

}
