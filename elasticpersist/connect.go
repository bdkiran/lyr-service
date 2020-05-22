package elasticpersist

import (
	"io"
	"io/ioutil"
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

//ES client to be used acrross the client
var es *elasticsearch.Client

//ConnectToEs creats a connection to elasticsearch
func ConnectToEs() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			//Connections address goes here
			"",
		},
	}
	var err error
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

//Lyric is the structure that represnts data store in elasticsearch
type Lyric struct {
	Artist     string `json:"artist"`
	Title      string `json:"title"`
	Lyric      string `json:"lyric"`
	LineNumber int    `json:"lineNumber"`
}

func getElasticSearchInfo() {
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)

	io.Copy(ioutil.Discard, res.Body)
}

//GetByID gets a a document by id
func GetByID(id string) {
	res, err := es.Get("song_lyrics", id)
	if err != nil {
		log.Printf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)

	io.Copy(ioutil.Discard, res.Body)
}

//GetLyricstBySongName gets documents by a songname
func GetLyricstBySongName(title string) []Lyric {
	return searchIndexedDocument(title, "title")
}

//GetLyricstByArtistName gets documents by an artist
func GetLyricstByArtistName(artist string) []Lyric {
	return searchIndexedDocument(artist, "artist")
}
