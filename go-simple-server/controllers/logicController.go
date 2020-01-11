package controller

import (
	"io/ioutil"
	"log"
	"net/http"
)
// DoMainRequest : retrieves main chunk of data
func DoMainRequest() []byte {
	newReq, anyError := http.NewRequest("GET", "http://api.pathofexile.com/public-stash-tabs", nil)
	//571783755-589642300-558214129-635578319-604141055
	if anyError != nil {
		log.Fatalln(anyError)
	}
	client := &http.Client{}
	q := newReq.URL.Query()
	q.Add("id", "571783755-589642300-558214129-635578319-604141055")
	newReq.URL.RawQuery = q.Encode()

	resp, anyError := client.Do(newReq)
	if anyError != nil {
		log.Fatalln(anyError)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	return body
}
