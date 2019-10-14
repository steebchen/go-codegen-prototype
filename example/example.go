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

	// default query API
	user, err := client.Post.FindOne.Where(
		Post.Title.Equals("welcome"),
	).Exec(ctx)
	log.Printf("user, err %+v %s", user, err)

	// advanced aggregation/order by query
	var result MyQuery1
	err = client.Post.Select.Name("MyQuery1").Fields(
		Post.Likes.Sum(),
	).GroupBy(
		Post.Title.Group(),
	).Into(&result).Exec(ctx)
	if err != nil {
		panic(err)
	}

	for _, item := range result {
		log.Printf("title: %s", item.Title)
		log.Printf("likes: %d", item.Likes)
	}
}
