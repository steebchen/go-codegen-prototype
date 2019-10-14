package main

import (
	"context"
	"log"

	. "go-codegen/example/photon"
)

func main() {
	client := NewClient()
	ctx := context.Background()

	// sql equivalent:
	// SELECT sum(likes), count(id)
	// FROM Posts
	// GROUP BY title

	var countResult CountQuery
	err := client.Post.Select.Name("CountQuery").Fields(
		Post.Count(),
	).Into(&countResult).Exec(ctx)

	if err != nil {
		panic(err)
	}

	log.Printf("count: %d", countResult[0].PostCount)
}
