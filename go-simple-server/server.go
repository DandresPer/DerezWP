package main

import (
	"encoding/json"
	"fmt"
	"github.com/Nuper/DerezWP/go-simple-server/controllers"
	"github.com/Nuper/DerezWP/go-simple-server/models"
)

func main() {

	myitems := controllers.FindItemsByName("Cold Penetration Support", rawData)

	var f models.FullResponse
	json.Unmarshal(rawData, &f)
	fmt.Println(f.NextChangeID)
	fmt.Println(f.Stashes[0].AccountName)
	fmt.Println(f.Stashes[0].Items[0].Verified)
	fmt.Println(len(myitems))
}