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
func FindItemsByName(name string, source []byte) []models.Item{
	rawData := DoMainRequest()
	println(string(rawData[0:300]))

	var f models.FullResponse
	json.Unmarshal(rawData, &f)
	itemArray := f.Stashes
	var filteredItems []models.Item
	for i := 0; i < len(itemArray); i++{
		for j := 0; j < len(itemArray[i].Items); i++{
			if(itemArray[i].Items[j].Name == name){
				filteredItems = append(filteredItems, itemArray[i].Items[j])
			}
		}
		fmt.Println(f.Stashes[i].AccountName)	
	}
	return filteredItems
}
