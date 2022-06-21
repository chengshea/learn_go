package test

import (
	"bytes"
	"context"
	"demo/com/cs/learn/tool"
	"encoding/json"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var (
	r map[string]interface{}
)

func queryAllES() {
	es := tool.OpenES()
	//version(es)

	for i, v := range []string{"Test 第三", "Test 第四"} {
		tool.EScreate(struct{ Title string }{Title: v}, "test", strconv.Itoa(i+3), es)

	}

	//queryES(es)

}

func version(es *elasticsearch.Client) {
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	log.Println("info:", res)
}

// Perform the search request.
func queryES(es *elasticsearch.Client) {

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"Title": "two",
			},
		},
	}

	r = tool.ESquery(query, "test", es)
	log.Println("data:", r)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

}

func createES(es *elasticsearch.Client, i int, title string) {
	data, err := json.Marshal(struct{ Title string }{Title: title})
	if err != nil {
		log.Fatalf("Error marshaling document: %s", err)
	}
	log.Println("data:", data)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      "test",
		DocumentID: strconv.Itoa(i + 1),
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}
