package main

import "fmt"

//Books 书籍
type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

func main() {
	//创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go语言教程", 6495407})
	//也可以使用key=>value
	fmt.Println(Books{title: "Go语言", author: "www.runoob.com", subject: "Go 语言教程", bookID: 6495407})
	//忽略的字段为0或空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	var Book1 Books
	Book1.title = "Go语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookID = 6495407

	fmt.Printf("Book 1 title: %s\n", Book1.title)
	fmt.Printf("Book 1 author: %s\n", Book1.author)
	fmt.Printf("Book 1 Subject: %s\n", Book1.subject)
	fmt.Printf("Book 1 bookID: %d\n", Book1.bookID)
}
