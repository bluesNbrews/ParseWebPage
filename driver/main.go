package main

import (
	"os"

	link "github.com/bluesNbrews/ParseWebPage/link"
	web "github.com/bluesNbrews/ParseWebPage/web"
)

func main() {

	//Retrieve URL string from input
	var enteredurl = string(os.Args[1])

	//Call the URL string and retrieve HTML content as io.Reader
	htmlcontent := web.Gethtml(enteredurl)

	//Parse htmlcontent and return an array of <a> tags found in it
	links, err := link.Parse(htmlcontent)
	if err != nil {
		panic(err)
	}

	//Read links array, add domain name to hrefs where missing, then place in new array
	newlinks := link.Fixlinks(links, enteredurl)

	//Query HREF URLs from the new array and display text and HTTP return code for each
	web.Getandprinturlstatus(newlinks)

}
