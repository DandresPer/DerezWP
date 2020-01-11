package main

import (
	"encoding/json"
	"fmt"
	"github.com/Nuper/DerezWP/go-simple-server/controllers"
	"github.com/Nuper/DerezWP/go-simple-server/models"
)

func main() {

body := controllers.DoMainRequest()
	println(string(body[0:300]))

	//println(len(string(body)))

	var f models.FullResponse
	json.Unmarshal(body, &f)
	fmt.Println(f.NextChangeID)
	fmt.Println(f.Stashes[0].AccountName)
	fmt.Println(f.Stashes[0].Items[0].Verified)

}
