package elasticpersist

import (
	"encoding/json"
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
	connectionURL, connectErr := utils.GetEnvVariableString("ES_CONNECTION_STRING")
	if connectErr != nil {
		logger.Error.Fatal(connectErr)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			connectionURL,
		},
	}
	var err error
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		logger.Error.Fatalf("Exiting due to error creating the client connection: %s", err)
	}

	//Ping the ES cluster before displaying this message
	logger.Info.Println("ES connection successful")
}

//Lyric is the structure that represnts data store in elasticsearch
type Lyric struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Lyric  string `json:"lyric"`
	//LyricDocID string  `json:"lyricDocID"`
	//LineNumber []int   `json:"lineNumber"`
	Upvotes []int64 `json:"upvotes"`
	DocID   string  `json:"docID"`
}

func checkElasticSearchDetails() {
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

//GetHealthOfCluster returns the status of the es cluster
func GetHealthOfCluster() (string, error) {
	res, err := es.Cluster.Health()
	if err != nil {
		logger.Warning.Printf("Error getting health status of cluster: %s", err)
		return "", err
	}

	defer res.Body.Close()

	var r map[string]interface{}

	log.Println(res)

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logger.Warning.Printf("Error parsing the response body: %s", err)
		return "", err
	}

	statusColor := r["status"]

	return statusColor.(string), nil
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
