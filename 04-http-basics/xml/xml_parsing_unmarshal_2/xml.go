package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// to store the name of XML element add a field named XMLName with the type xml.Name

// To store the attribute of an XML element, define a field with the same name as that attribute and use the struct tag `xml:"<name>,attr"`, where <name> is the name of the XML attribute.

// to get to an XML element directly without specifying the tree structure, use the struct tag `xml:"a>b>c"`, where a and b are the intermediate elements and c is the node that you want to get to
type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	// read XML file and conver to bytes 
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post.XMLName.Local)
	fmt.Println(post.Id)
	fmt.Println(post.Content)
	fmt.Println(post.Author)
	fmt.Println(post.Xml)
	fmt.Println(post.Author.Id)
	fmt.Println(post.Author.Name)
	fmt.Println(post.Comments)
	fmt.Println(post.Comments[0].Id)
	fmt.Println(post.Comments[0].Content)
	fmt.Println(post.Comments[0].Author)
	fmt.Println(post.Comments[1].Id)
	fmt.Println(post.Comments[1].Content)
	fmt.Println(post.Comments[1].Author)
}
