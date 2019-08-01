/* Example of storing data in memory 

*/

package main 

import (
	"fmt"
)

type Post struct {
	Id int
	Content string 
	Author string
}

// storing post in data mean maping it to a key
var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

// store functions stores a pointer to the post 
func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post.Content)
	}

}