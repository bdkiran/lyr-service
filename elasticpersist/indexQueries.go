package elasticpersist

import (
	"context"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

const subscriptionIndexName = "subscriptions"
const lyricsSubmissionIndexName = "subscriptions"

func indexEmails(indexBody string) {
	req := esapi.IndexRequest{
		Index:   subscriptionIndexName,
		Body:    strings.NewReader(indexBody),
		Refresh: "true",
	}

	req.Do(context.Background(), es)
}

func indexLyrics() {

}
