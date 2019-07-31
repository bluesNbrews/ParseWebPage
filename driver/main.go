package main

import (

	"os"
	"fmt"
	"time"

	"github.com/bluesNbrews/ParseWebPage/link"
	"github.com/bluesNbrews/ParseWebPage/web"

)

func main() {

	//Retrieve URL string from input
	var enteredurl = string(os.Args[1])

	//String it for input to keep program running
	var input string
	
	//Channel for status code
	var c chan int = make(chan int)

	//Call the URL string and retrieve HTML content as io.Reader
	htmlcontent := web.Gethtml(enteredurl)

	//Parse htmlcontent and return an array of <a> tags found in it
	links, err := link.Parse(htmlcontent)
	if err != nil {
		panic(err)
	}

	//Read links array, add domain name to hrefs where missing, then place in new array
	newlinks := link.Fixlinks(links, enteredurl)

	for i := 0; i < len(links); i++ {
		fmt.Println(i, links[i])
	}

	//Process each link concurrently to get http status code and assign to link and print
	//The sleep is used temporarily to prevent too many http requests at one time
	for i := 0; i < len(newlinks); i++ {
		
		go web.GetUrlStatus(newlinks[i], c)
		go web.UpdateAndPrint(newlinks[i], c)
		time.Sleep(125 * time.Millisecond)
		
	}

	//Looks for user input to keep program running while http requests are processing
	fmt.Scanln(&input)
}
