package main

import (
	"context"
	"fmt"

	"github.com/ObsidianCat/the-3-ages-of-go-concurrency-talk/internal"
)

func main() {
	// pretty print createUsersRowsWithChannels func result map
	fmt.Println("createUsersRowsWithChannels result:")
	result := createUsersRowsWithChannels(context.Background())
	prettyPrint(result)

	fmt.Println("createUsersRowsWithSemaphore result:")
	// pretty print createUsersRowsWithSemaphore func result map
	result, _ = createUsersRowsWithSemaphore(context.Background())
	prettyPrint(result)

}

func prettyPrint(printable map[string]internal.UserRow) {
	for k, v := range printable {
		// pretty print every user row in the map
		fmt.Printf("%v: %v\n", k, v)
	}
}
