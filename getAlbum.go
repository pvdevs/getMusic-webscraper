package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type Album struct {
	Title	string
	Artist	string
	Rating	string
	Year	string
	Imgurl	string
}

type JsonResponse struct {
	Count 		int
	Previous 	int
	Next 		int
	Results 	Results
}

type Results struct {
	Category 	Category
	List 		[]List
}

type Category struct {
	Header				string
	Id					int
	Name				string
	Bio					string
	MobileHeader		string
	SocialTitle			string
	SocialDescription	string
	SocialImage			string
	Url					string
}

type List struct {

}

func main() {
	resp, error := http.Get("https://pitchfork.com/api/v2/search/?genre=experimental&types=reviews&hierarchy=sections%2Freviews%2Falbums%2Cchannels%2Freviews%2Falbums&sort=publishdate%20desc%2Cposition%20asc&size=12&start=2820")

	if error != nil {
		// Handle error
	}

	fmt.Println(reflect.TypeOf(resp))

	b, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		panic(err)
	}

	var jsonResponse JsonResponse

	er := json.Unmarshal([]byte(b), &jsonResponse)

	if er != nil {
		log.Fatalf("Unable to marshal JSON due to %s", er)
	}

	fmt.Printf("%s", jsonResponse.Results.Category.Name)

	defer resp.Body.Close()
	//body, err := io.ReadAll(resp.Body)
}
