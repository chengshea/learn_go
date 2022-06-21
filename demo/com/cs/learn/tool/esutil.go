package tool

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var (
	es *elasticsearch.Client
	r  map[string]interface{}
)

func initES() *elasticsearch.Client {
	//cert, _ := ioutil.ReadFile(*cacert)
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			// "https://localhost:9201",
		},
		// Username: "foo",
		// Password: "bar",
		//  CACert: cert,
	}
	clinet, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalln("init es err:", err)
	}
	return clinet
}

func OpenES() *elasticsearch.Client {
	if es == nil {
		log.Println("init con ES")
		es = initES()
	}
	return es
}

func EScreate(a any, index string, id string, es *elasticsearch.Client) {
	if a == nil {
		log.Fatalf("Error body is nil")
	}

	data, err := json.Marshal(a)
	if err != nil {
		log.Fatalf("Error marshaling document: %s", err)
	}

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
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
		log.Printf("[%s] Error indexing document ID=%s", res.Status(), id)
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

func ESquery(query map[string]interface{}, index string, es *elasticsearch.Client) map[string]interface{} {

	if query == nil || index == "" {
		log.Fatalf("Error query:%s,index:%s is nil", query, index)
	}
	bf := bytes.Buffer{}
	err := json.NewEncoder(&bf).Encode(query)

	if err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&bf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			//Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	defer res.Body.Close()
	return r
}
