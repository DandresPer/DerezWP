package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"github.com/Nuper/DerezWP/go-simple-server/models"
	"time"
)

// DoMainRequest : retrieves SOURCE data
func DoMainRequest(nextChangeID string) []byte {
	
	newReq, anyError := http.NewRequest("GET", "http://api.pathofexile.com/public-stash-tabs", nil)
	//571783755-589642300-558214129-635578319-604141055
	if anyError != nil {
		log.Fatalln(anyError)
	}
	client := &http.Client{}
	q := newReq.URL.Query()

	if(nextChangeID != "") {
		q.Add("id", nextChangeID)
	}
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
var currentNextChangeID = ""
var lastNextChangeID = "lastNextChangeID"
// FindItemsByName : Finds items by name in the SOURCE
func FindItemsByName(name string) []models.Item{
	var FilteredItems []models.Item
	start := time.Now()
	println(start.String())
	for currentNextChangeID != lastNextChangeID 	{
		lastNextChangeID = currentNextChangeID
		rawData := DoMainRequest(currentNextChangeID)
		var f models.FullResponse
		json.Unmarshal(rawData, &f)
		currentNextChangeID = f.NextChangeID
		itemArray := f.Stashes
		
		for i := 0; i < len(itemArray); i++{
			for j := 0; j < len(itemArray[i].Items); j++{
				if(itemArray[i].Items[j].TypeLine == name){
					println(currentNextChangeID);
					println("MATCH FOUND! " + itemArray[i].Items[j].TypeLine)
					FilteredItems = append(FilteredItems, itemArray[i].Items[j])
				}
			}
		}
		
	}
	elapsed := time.Since(start)
	println(elapsed.String())
	println("EOF");
	return FilteredItems
}
