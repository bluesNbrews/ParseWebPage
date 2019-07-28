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

	//Parse htmlcontent and return a slice of <a> tags found in it
	links, err := link.Parse(htmlcontent)
	if err != nil {
		panic(err)
	}

	//Retrieve the HREFS from the slice, remove duplicates,
	//append domain name where missing, then place in new slice
	uniqueurls := link.Gethrefs(links, enteredurl)

	//Query URLs from the new slice and display HTTP return code for each
	web.Getandprinturlstatus(uniqueurls)

}
