package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type AlbumData struct {
	Title	string
	Artist	string
	Genre	string
	Rating	string
	Year	int
	Imgurl	string
}

type JsonResponse struct {
	Results 	Results
	Count		int
}

type Results struct {
	List	[]List
}

type List struct {
	Tombstone	Tombstone
	Genres		[]Genres
}

type Genres struct {
	DisplayName	string	`json:"display_name"`
}

type Tombstone struct {
	Albums	[]Albums
}

type Albums struct {
	Album 	Album
	Rating	Rating
}

type Rating struct {
	Rating string
}

type Album struct {
	Artists		[]Artists
    DisplayName string	`json:"display_name"`
	ReleaseYear	int		`json:"release_year"`
	Photos		Photos

}

type Artists struct {
	DisplayName	string	`json:"display_name"`
}

type Photos struct {
	Tout 	Tout
	Title	string
	AltText	string
}
type Tout struct {
	Sizes Sizes
}
type Sizes struct {
	Standard string
}


func getAllAlbums(url string) int {
	resp, error := http.Get(url)
	
	if error != nil {
		// Handle error
	}


	b, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		panic(err)
	}

	var jsonResponse JsonResponse

	er := json.Unmarshal([]byte(b), &jsonResponse)

	if er != nil {
		log.Fatalf("Unable to marshal JSON due to %s", er)
	}

	count := jsonResponse.Count

	return count
}

func getRandomAlbumPosition(url string) string {
	count := getAllAlbums(url)

	println("total of albums:", count)

	randomCount := rand.Intn(count +1)

	println("this generated album position:", randomCount)

	randomCountStr := strconv.Itoa(randomCount)

	return randomCountStr
}

func main() {
	rand.Seed(time.Now().UnixNano())

	position := getRandomAlbumPosition("https://pitchfork.com/api/v2/search/?genre=electronic&genre=experimental&genre=jazz&types=reviews&hierarchy=sections%2Freviews%2Falbums%2Cchannels%2Freviews%2Falbums&sort=publishdate%20desc%2Cposition%20asc&size=1&start=12")

	resp, error := http.Get("https://pitchfork.com/api/v2/search/?genre=electronic&genre=experimental&genre=jazz&types=reviews&hierarchy=sections%2Freviews%2Falbums%2Cchannels%2Freviews%2Falbums&sort=publishdate%20desc%2Cposition%20asc&size=1&start="+position)

	if error != nil {
		// Handle error
	}

	b, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		panic(err)
	}

	var jsonResponse JsonResponse

	er := json.Unmarshal([]byte(b), &jsonResponse)

	if er != nil {
		log.Fatalf("Unable to marshal JSON due to %s", er)
	}

	var albumData AlbumData

	// Must loop through each artist and store their respective names
	albumData.Artist = jsonResponse.Results.List[0].Tombstone.Albums[0].Album.Artists[0].DisplayName

	albumData.Title = jsonResponse.Results.List[0].Tombstone.Albums[0].Album.DisplayName

	// Must loop through each genre and store their respective names
	albumData.Genre = jsonResponse.Results.List[0].Genres[0].DisplayName

	albumData.Year = jsonResponse.Results.List[0].Tombstone.Albums[0].Album.ReleaseYear

	albumData.Rating = jsonResponse.Results.List[0].Tombstone.Albums[0].Rating.Rating

	albumData.Imgurl = jsonResponse.Results.List[0].Tombstone.Albums[0].Album.Photos.Tout.Sizes.Standard

	fmt.Println(albumData)

	defer resp.Body.Close()
}
