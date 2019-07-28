package web

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
func Getandprinturlstatus(uniqueurls []string) {
	for i := 0; i < len(uniqueurls); i++ {
		resp, err := http.Get(uniqueurls[i])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-3s- %-60s| %-4s|\n", strconv.Itoa(i+1), uniqueurls[i], strconv.Itoa(resp.StatusCode))
	}
}
