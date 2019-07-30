package web

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bluesNbrews/ParseWebPage/link"
)

//Gethtml makes a GET call to a URL and returns the HTML body
func Gethtml(enteredurl string) io.Reader {
	
	resp, err := http.Get(enteredurl)
	
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	htmlcontent := bytes.NewReader(body)

	return htmlcontent

}

//Get http respone code for each link passed in. Use channel to pass back response code, 
//which will be concurrently assigned and printed via UpdateAndPrint
func GetUrlStatus(newlinks link.Link, c chan int) {

	if newlinks.Href != ""{
		
		resp, err := http.Get(newlinks.Href)
		
		if err != nil {
			log.Fatal(err)
		}

		c <- resp.StatusCode

	} 
}

func UpdateAndPrint(newlinks link.Link, c chan int){
	
	//Assign status code via channel
	newlinks.Code = <- c

	//Print table rows
	fmt.Printf("| %-80s | %-40s | %-5d|\n", newlinks.Href, newlinks.Text, newlinks.Code)
}
