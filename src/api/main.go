package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mauryparra/elastic-go/src/api/domain"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

const (
	elasticIndexName = "items"
	elasticTypeName  = "item"
)

var (
	elasticClient *elastic.Client
)

func main() {
	var err error
	// Create Elastic client and wait for Elasticsearch to be ready
	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("http://localhost:9200"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			// Retry every 3 seconds
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}

	fmt.Println("Elastic detected")

	// Start HTTP server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/items/:ID", getItemFromElastic)
	r.POST("/items/:ID", createItemElastic)
	r.GET("/search/items", searchEndpoint)

	if err = r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func createItemElastic(c *gin.Context) {

	item := domain.Item{
		ID: c.Param("ID"),
	}

	if err := item.Get(); err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to create item")
		return
	}

	bulk := elasticClient.
		Bulk().
		Index(elasticIndexName).
		Type(elasticTypeName)

	bulk.Add(elastic.NewBulkIndexRequest().Id(item.ID).Doc(item))

	if _, err := bulk.Do(c.Request.Context()); err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to create item")
		return
	}
	c.String(http.StatusOK, "Item added")
}

func getItemFromElastic(c *gin.Context) {
	getItem, err := elasticClient.Get().
		Index(elasticIndexName).
		Type(elasticTypeName).
		Id(c.Param("ID")).
		Do(c)

	if err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to create item")
		return
	}

	if getItem.Found {
		c.JSON(http.StatusAccepted, getItem.Source)
		return
	}

	errorResponse(c, http.StatusNotFound, "Item not found")
	return
}

func searchEndpoint(c *gin.Context) {
	// Parse request
	query := c.Query("query")
	if query == "" {
		errorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}
	skip := 0
	take := 10
	if i, err := strconv.Atoi(c.Query("skip")); err == nil {
		skip = i
	}
	if i, err := strconv.Atoi(c.Query("take")); err == nil {
		take = i
	}
	// Perform search
	esQuery := elastic.NewMultiMatchQuery(query, "title", "subtitle").
		Fuzziness("2").
		MinimumShouldMatch("2")
	result, err := elasticClient.Search().
		Index(elasticIndexName).
		Query(esQuery).
		From(skip).Size(take).
		Do(c.Request.Context())
	if err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}
	c.JSON(http.StatusOK, result.Hits.Hits)
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
