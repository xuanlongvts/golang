package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
}

func main() {
	log.SetFlags(0)

	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second * 5,
			}).DialContext,
			TLSClientConfig:       &tls.Config{MinVersion: tls.VersionTLS12},
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 5,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal("err elasticsearch.NewClient: ---> ", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatal("err es.Info: ---> ", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.Status())
	}

	fmt.Println(res.String())

	// Create an index
	createIndex(es)

	// Index sample user data
	users := []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Location: "New York"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Location: "London"},
		{ID: 3, Name: "Michael Johnson", Email: "michael@example.com", Location: "San Francisco"},
	}
	for _, user := range users {
		indexDocument(es, user)
	}

	// Perform a basic search
	query := `{
		"query": {
			"match": {
				"name": "john"
			}
		}
	}`

	search(es, query)
}

func createIndex(es *elasticsearch.Client) {
	indexName := "users"

	req := esapi.IndicesCreateRequest{
		Index: indexName,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error creating index: %s", res.Status())
	} else {
		fmt.Println("Index created successfully")
	}
}

func indexDocument(es *elasticsearch.Client, user User) {
	indexName := "users"

	docID := fmt.Sprintf("%d", user.ID)
	doc := fmt.Sprintf(`{
		"id": %d,
		"name": "%s",
		"email": "%s",
		"location": "%s"
	}`, user.ID, user.Name, user.Email, user.Location)

	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: docID,
		Body:       strings.NewReader(doc),
		Refresh:    "true",
		Pretty:     true,
		ErrorTrace: true,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatal("err: req.Do: ", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error indexing document: %s", res.Status())
	} else {
		fmt.Printf("Document indexed successfully: ID=%s\n", docID)
	}
}

func search(es *elasticsearch.Client, query string) {
	indexName := "users"

	esContext := es.Search.WithContext(context.Background())
	esIndex := es.Search.WithIndex(indexName)
	esBody := es.Search.WithBody(strings.NewReader(query))
	esTrackTotalHits := es.Search.WithTrackTotalHits(true)
	esPretty := es.Search.WithPretty()
	res, err := es.Search(esContext, esIndex, esBody, esTrackTotalHits, esPretty)
	if err != nil {
		log.Fatal("err: ", err.Error())
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error searching document: %s", res.Status())
	}

	fmt.Println("result: ", res.String())
}
