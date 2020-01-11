package main

import (
	"encoding/json"
	"fmt"
	"github.com/Nuper/DerezWP/go-simple-server/controllers"
)

// func main() {
// 	http.HandleFunc("/", homeHandler)
// 	http.ListenAndServe(":3000", nil)
// }

// func homeHandler(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("Hello World"))
// }

func main() {

body := controller.DoMainRequest()
	println(string(body[0:300]))

	type Requirement struct {
		Name string
		//Values
		DisplayMode int
	}

	type Property struct {
		Name string
		//Values
		DisplayMode int
	}

	type Item struct {
		Verified   bool
		W          int
		H          int
		Icon       string
		League     string
		ID         string
		Name       string
		TypeLine   string
		Identified string
		Ilvl       string
		Properties []Property
		// //ExplicitiMods
		// //flavourText
		// //frameType
		StackSize    int
		MaxStackSize int
		ArtFileName  string
		//extended
		X           int
		Y           int
		InventoryID string
	}

	type Stash struct {
		ID                string
		Public            bool
		AccountName       string
		LastCharacterName string
		Stash             string
		StashType         string
		League            string
		Items             []Item
	}

	type FullResponse struct {
		NextChangeID string `json:"next_change_id"`
		Stashes      []Stash
	}
	//println(len(string(body)))

	var f FullResponse
	json.Unmarshal(body, &f)
	fmt.Println(f.NextChangeID)
	fmt.Println(f.Stashes[0].AccountName)
	fmt.Println(f.Stashes[0].Items[0].Verified)

}
