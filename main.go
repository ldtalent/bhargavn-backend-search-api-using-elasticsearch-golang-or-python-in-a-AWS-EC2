package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/olivere/elastic/v7"
)

type SearchRequest struct {
	// Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var searchRequest SearchRequest
	err := json.NewDecoder(r.Body).Decode(&searchRequest)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetBasicAuth("bhargavnath", "nath123"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		http.Error(w, "Failed to create Elasticsearch client", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	boolQuery := elastic.NewBoolQuery()

	if searchRequest.FirstName != "" {
		boolQuery = boolQuery.Must(elastic.NewMatchQuery("first_name", searchRequest.FirstName))
	}
	if searchRequest.LastName != "" {
		boolQuery = boolQuery.Must(elastic.NewMatchQuery("last_name", searchRequest.LastName))
	}

	searchResult, err := client.Search().
		Index("students").
		Query(boolQuery).
		Do(ctx)
	if err != nil {
		http.Error(w, "Failed to execute search query", http.StatusInternalServerError)
		return
	}

	var searchResponse []map[string]interface{}
	for _, hit := range searchResult.Hits.Hits {
		var result map[string]interface{}
		err := json.Unmarshal(hit.Source, &result)
		if err != nil {
			http.Error(w, "Failed to unmarshal search result", http.StatusInternalServerError)
			return
		}
		searchResponse = append(searchResponse, result)
	}

	if len(searchResponse) == 0 {
		http.Error(w, "No results found for the given search criteria", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searchResponse)
}

func main() {
	http.HandleFunc("/search", searchHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
