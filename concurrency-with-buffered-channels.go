package main

import (
	"context"

	"github.com/ObsidianCat/the-3-ages-of-go-concurrency-talk/domain"
)

func createUsersRowsWithChannels(ctx context.Context) map[string]domain.UserRow {
	// Fetch data
	profilesByUserID := domain.FetchStaffProfiles(ctx)

	// 1 Define variables
	numOfProfiles := len(profilesByUserID)

	// Replacement of userRowMtx and concurrentgroup
	userRowsChan := make(chan domain.UserRow, numOfProfiles)
	concurrencyLimiter := make(chan struct{}, 100)

	// 2 Create user rows concurrenty
	for _, userProfile := range profilesByUserID {
		// Acquire a "slot" in the semaphore, stop and wait if not available
		concurrencyLimiter <- struct{}{}

		go func(profile domain.UserProfile, userRowsChan chan<- domain.UserRow, concurrencyLimiter <-chan struct{}) {
			// Release the "slot" in the semaphore
			defer func() {
				<-concurrencyLimiter
			}()

			row := domain.NewUserRow(ctx, profile)

			// Send data to the channel
			userRowsChan <- row
		}(userProfile, userRowsChan, concurrencyLimiter)
	}

	// 3 Convert channel messages to the map of user rows and return them
	rows := make(map[string]domain.UserRow, numOfProfiles)
	for i := 0; i < numOfProfiles; i++ {
		row := <-userRowsChan
		rows[row.ID] = row
	}

	return rows
}
