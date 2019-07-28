package web

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/bluesNbrews/ParseWebPage/link"
)

//Gethtml makes a GET call to a URL and returns the HTML body
func Gethtml(enteredurl string) io.Reader {
	//Make GET request to the entered URL and store HTML in variable htmlcontent
	resp, err := http.Get(enteredurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	htmlcontent := bytes.NewReader(body)

	return htmlcontent
}

//Getandprinturlstatus makes a GET call to the passed URLs and prints status for each
func Getandprinturlstatus(newlinks []link.Link) {

	//Print table header
	fmt.Printf("%-3s| %-60s | %-30s | %-16s|\n", "#", "URL LINKS", "URL TEXT", "CODE")

	//Iterate over each item in newlinks array, if not previously check then get HTTP code, and display output
	for i := 0; i < len(newlinks); i++ {

		//Check if item was previously checked
		prevchecked := previouslychecked(newlinks[i].Href, newlinks, i)

		//If item not previously checked, then get HTTP return code, else set return code to (prev. checked)
		var returncode string

		if !(prevchecked) {
			resp, err := http.Get(newlinks[i].Href)
			if err != nil {
				log.Fatal(err)
			}
			returncode = strconv.Itoa(resp.StatusCode)
		} else {
			returncode = "(prev. checked)"
		}

		//Print table rows
		fmt.Printf("%-3s| %-60s | %-30s | %-16s|\n", strconv.Itoa(i+1), newlinks[i].Href, newlinks[i].Text, returncode)
	}
}

func previouslychecked(href string, links []link.Link, i int) bool {
	for j := 0; j < i; j++ {
		if links[j].Href == href {
			return true
		}
	}
	return false
}
