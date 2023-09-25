package main

import (
	"context"
	"fmt"

	"github.com/ObsidianCat/the-3-ages-of-go-concurrency-talk/domain"
)

func main() {
	// pretty print createUsersRowsWithChannels func result map
	prettyPrint(createUsersRowsWithChannels(context.Background()))

}

func prettyPrint(printable map[string]domain.UserRow) {
	for k, v := range printable {
		// pretty print every property with it name
		fmt.Printf("%v: %v\n", k, v)
	}
}
