/* 
If need to stream data or big XML files, use Decoder instead of Unmarshal 
*/

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

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
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// create decoder
	decoder := xml.NewDecoder(xmlFile)

	// iterate through XML data in decoder
	for {
		t, err := decoder.Token() // get token from decoder at each iteration

		// when run out of tokens, the err will not be nil 
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se) // decode XML data into struct
				fmt.Println(comment)
			}
		}
	}
}
