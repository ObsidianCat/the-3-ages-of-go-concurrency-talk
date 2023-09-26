package main

import (
	"context"

	"github.com/ObsidianCat/the-3-ages-of-go-concurrency-talk/domain"
	"golang.org/x/sync/semaphore"
)

func createUsersRowsWithSemaphore(ctx context.Context) (map[string]domain.UserRow, error) {
	// 1 Fetch data
	profilesByUserID := domain.FetchStaffProfiles(ctx)

	// 2 Define variables
	numOfProfiles := len(profilesByUserID)
	rowsChan := make(chan domain.UserRow, numOfProfiles)
	concurrencyLimiter := semaphore.NewWeighted(2)

	// 3 Create user rows concurrenty
	for _, userProfile := range profilesByUserID {
		// Acquire a "token" from semaphore, block if none available.
		if err := concurrencyLimiter.Acquire(ctx, 1); err != nil {
			return nil, err
		}

		go func(profile domain.UserProfile, rowsChan chan<- domain.UserRow, concurrencyLimiter *semaphore.Weighted) {
			// Release the "token" back to semaphore
			defer concurrencyLimiter.Release(1)

			// Create user row including the data from the profile
			// and fetchign more data by calling other services
			// with user ID from the profile
			row := domain.NewUserRow(ctx, profile)
			rowsChan <- row
		}(userProfile, rowsChan, concurrencyLimiter)
	}

	// 4 Convert channel messages to the map of user rows and return them
	rows := make(map[string]domain.UserRow, numOfProfiles)
	for i := 0; i < numOfProfiles; i++ {
		userRow := <-rowsChan
		rows[userRow.ID] = userRow
	}

	return rows, nil
}
