package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
)

// func main() {
// 	http.HandleFunc("/", homeHandler)
// 	http.ListenAndServe(":3000", nil)
// }

// func homeHandler(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("Hello World"))
// }

func main() {
	resp, err := http.Get("http://api.pathofexile.com/public-stash-tabs/?id=571783755-589642300-558214129-635578319-604141055")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}
		
		println(string(body[0:300]))
		// type Stash struct {

		//     id string
		//     public bool
		//     accountName string
		//     lastCharacterName string
		//     stash string
		//     stashType string
		//     league string
		// }
		type FullResponse struct {

			next_change_id string

		}
			println(len(string(body)))

	var f FullResponse
			erro := json.Unmarshal(body, &f)
			//b := []byte('{"Name":"Bob"}')
			if (erro == nil){
				fmt.Printf("%+v\n", f)
			} else { println(erro)}
			
} 

