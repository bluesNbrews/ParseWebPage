package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/bluesNbrews/ParseWebPage/link"
)

var exampleHTML = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {

	//Retrieve URL to query
	var fullpath = string(os.Args[1])

	//Retrieve domain name from URL (to be used later for appending to URL paths)
	zp := regexp.MustCompile(`/`)
	var temp = zp.Split(fullpath, -1)
	var domainname = temp[0] + "//" + temp[2]

	//Make GET request to the entered URL and store HTML in variable r
	resp, err := http.Get(fullpath)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	r := bytes.NewReader(body)

	//Parse will take in an HTML document (r) and will return a slice of links parsed from it.
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(links); i++ {
		var path = links[i].Href
		if strings.HasPrefix(path, "/") {
			path = domainname + path
		}
		fmt.Println(path)
	}

}
