package main

import (
	"fmt"
)

type Post struct {
	body string
	publishDate int64
	next *Post
}

type Feed struct {
	length int
	start *Post
	end *Post
}

// 添加一个元素
func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
	} else {
		lastPost := f.end
		lastPost.next = newPost
	}
	f.end = newPost
	f.length++
}
// 删除一个元素
func (f *Feed) Remove(publishDate int64) bool {
	if f.length == 0 {
		return false
	}

	var previousPost *Post
	currentPost := f.start
	for currentPost.publishDate != publishDate {
		if currentPost.next == nil {
			return false
		}
		previousPost = currentPost
		currentPost = currentPost.next
	}
	previousPost.next = currentPost.next
	f.length--
	return true
}

// 插入一个元素
func (f *Feed) Insert(post *Post) bool {
	if f.length == 0 {
		// 只有一个元素的时候
		f.start = post
		f.end = post
		return true
	}
	var previousPost *Post
	currentPost := f.start
	for currentPost !=nil && currentPost.publishDate < post.publishDate {
		previousPost = currentPost
		currentPost = previousPost.next
	}

	previousPost.next = post
	post.next = currentPost
	f.length++
	return true
}

func main() {
	f := &Feed{}
	p1 := Post{
		body:        "lorem ipsum",
		publishDate:1,
	}
	f.Append(&p1)
	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)
	fmt.Printf("Second: %v\n", f.start.next)
	p2 := Post{
		body: "Dolor sit amet",
		publishDate: 2,
	}
	f.Append(&p2)
	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)
	fmt.Printf("Second: %v\n", f.start.next)
	f.Append(&Post{
		body:        "the three post",
		publishDate: 3,
	})
	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)

	f.Insert(&Post{
		body:        "insert data to list",
		publishDate: 4,
		next:        nil,
	})

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)
}