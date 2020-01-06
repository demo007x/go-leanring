package main

import "fmt"

type author struct {
	firtName string
	lastName string
	bio      string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firtName, a.lastName)
}

type post struct {
	title string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title:", p.title)
	fmt.Println("Content:", p.content)
	fmt.Println("Author:", p.author.fullName())
	fmt.Println("Bio:", p.author.bio)
}

type website struct {
	post []post
}

func (w website) contents () {
	fmt.Println("Contents of Website")
	for _,v := range w.post {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		firtName: "Naveen",
		lastName: "Ramanathan",
		bio:      "Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post2 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post3 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	w := website{post:[]post{post1, post2, post3}}

	w.contents()
}
