package main

import (
	"context"

	"github.com/ObsidianCat/the-3-ages-of-go-concurrency-talk/domain"
)

func createUsersRowsWithChannels(ctx context.Context) (map[string]userRow, error) {
	// [Code omitted]

	profiles, err := domain.FetchStaffProfiles(ctx)
	if err != nil {
		return nil, err
	}
	profilesByUserID := domain.SliceToMapById(profiles)

	// 1 Define variables
	numOfProfiles := len(profilesByUserID)

	// Replacement of userRowMtx and concurrentgroup
	userRowsChan := make(chan domain.UserRow, numOfProfiles)
	concurrencyLimiter := make(chan struct{}, 100)

	// 2 Create user rows concurrenty
	for _, userProfile := range profilesByUserID {
		// Acquire a "slot" in the semaphore, stop and wait if not available
		concurrencyLimiter <- struct{}{}

		go func(profile *staffprofileproto.Profile, userRowsChan chan<- domain.UserRow, concurrencyLimiter <-chan struct{}) {
			// Release the "slot" in the semaphore
			defer func() {
				<-concurrencyLimiter
			}()

			row := newUserRow(ctx, profile, rolesByUserID, policiesByUserID, teamByUserID, collectiveByUserID)

			// Send data to the channel
			userRowsChan <- row
		}(userProfile, userRowsChan, concurrencyLimiter)
	}

	// 3 Convert channel messages to the map of user rows and return them
	rows := make(map[string]domain.UserRow, numOfProfiles)
	for i := 0; i < numOfProfiles; i++ {
		userRow := <-userRowsChan
		rows[userRow.id] = user
	}

	return rows, nil
}
