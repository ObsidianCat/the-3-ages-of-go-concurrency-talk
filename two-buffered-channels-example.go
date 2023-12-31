package main

import (
	"context"

	"github.com/ObsidianCat/the-3-ages-of-go-concurrency-talk/internal"
)

func createUsersRowsWithChannels(ctx context.Context) map[string]internal.UserRow {
	// 1 Fetch data
	profilesByUserID := internal.FetchStaffProfiles(ctx)

	// 2 Define variables
	numOfProfiles := len(profilesByUserID)

	// Replacement of userRowMtx and concurrentgroup
	rowsChan := make(chan internal.UserRow, numOfProfiles)
	concurrencyLimiter := make(chan struct{}, 2)

	// 3 Create user rows concurrenty
	for _, userProfile := range profilesByUserID {
		// Acquire a "slot" in the semaphore, stop and wait if not available
		concurrencyLimiter <- struct{}{}

		go func(profile internal.UserProfile, userRowsChan chan<- internal.UserRow, concurrencyLimiter <-chan struct{}) {
			// Release the "slot" in the semaphore
			defer func() {
				<-concurrencyLimiter
			}()

			// Create user row including the data from the profile
			// and fetchign more data by calling other services
			// with user ID from the profile
			row := internal.NewUserRow(ctx, profile)

			// Send data to the channel
			userRowsChan <- row
		}(userProfile, rowsChan, concurrencyLimiter)
	}

	// 4 Convert channel messages to the map of user rows and return them
	rows := make(map[string]internal.UserRow, numOfProfiles)
	for i := 0; i < numOfProfiles; i++ {
		userRow := <-rowsChan
		rows[userRow.ID] = userRow
	}

	return rows
}
