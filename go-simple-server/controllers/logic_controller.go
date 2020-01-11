package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/Nuper/DerezWP/go-simple-server/models"
)
// DoMainRequest : retrieves SOURCE data
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
// FindItemsByName : Finds items by name in the SOURCE
func FindItemsByName() []models.Item{
	rawData := DoMainRequest()
	println(string(rawData[0:300]))

	var f models.FullResponse
	json.Unmarshal(rawData, &f)
	var g []models.Item
	for i := 0; i < len(f.Stashes); i++{
		fmt.Println(f.Stashes[i])
	}
	return g
}