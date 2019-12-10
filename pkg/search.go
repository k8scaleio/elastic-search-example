package pkg

import (
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"log"
)

func SearchDocuments(id, query string) [] interface{} {
	es, err := elasticsearch7.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)


	return nil
}
