package elasticpersist

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/bdkiran/lyr-service/utils"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

//Initilize variable to access project logger,
//this initialization can be used accoss the whole package
var logger = utils.NewLogger()

//ES client to be used acrross the client
var es *elasticsearch.Client

//ConnectToEs creats a connection to elasticsearch
func ConnectToEs() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	var err error
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		logger.Error.Fatalf("Exiting due to error creating the client connection: %s", err)
	}

	//Get information about es cluster
	//when this is removed, i/o timout occurs
	res, err := es.Info()
	if err != nil {
		logger.Error.Fatalf("Exiting due to error getting response from Elasticsearch cluster: %s", err)
	}

	defer res.Body.Close()
	logger.Info.Print(res)

	io.Copy(ioutil.Discard, res.Body)
}

//Lyric is the structure that represnts data store in elasticsearch
type Lyric struct {
	Artist     string `json:"artist"`
	Title      string `json:"title"`
	Lyric      string `json:"lyric"`
	LineNumber int    `json:"lineNumber"`
}

//getByID gets a a document by id. Not sure if this will be used.
func getByID(id string) {
	res, err := es.Get("song_lyrics", id)
	if err != nil {
		logger.Warning.Printf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)

	io.Copy(ioutil.Discard, res.Body)
}
