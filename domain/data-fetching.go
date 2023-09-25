package domain

import (
	"context"
	"time"
)

func FetchStaffProfiles(ctx context.Context) ([]userProfile, error) {
	// Simulate a 0.5 second wait to fetch data
	time.Sleep(500 * time.Millisecond)

	// Return a mock result
	return []userProfile{
		{ID: "e3f6fbd6-5bac-11ee-8c99-0242ac120002", Name: "John Doe", Email: "john@xample.com"},
		{ID: "e3f6ff46-5bac-11ee-8c99-0242ac120002", Name: "Jane Smith", Email: "jane@example.com"},
	}, nil
}

func FetchStaffProfiles(ctx context.Context) ([]userProfile, error) {
	// Simulate a 0.5 second wait to fetch data
	time.Sleep(500 * time.Millisecond)

	// Return a mock result
	return []userProfile{
		{ID: "e3f6fbd6-5bac-11ee-8c99-0242ac120002", Name: "John Doe", Email: "john@xample.com"},
		{ID: "e3f6ff46-5bac-11ee-8c99-0242ac120002", Name: "Jane Smith", Email: "jane@example.com"},
	}, nil
}

func FetchUserRolesByID(ctx context.Context) ([]userProfile, error) {
	// Simulate a 0.5 second wait to fetch data
	time.Sleep(500 * time.Millisecond)

	// Return a mock result
	return []userProfile{
		{ID: "e3f6fbd6-5bac-11ee-8c99-0242ac120002", Name: "John Doe", Email: "john@xample.com"},
		{ID: "e3f6ff46-5bac-11ee-8c99-0242ac120002", Name: "Jane Smith", Email: "jane@example.com"},
	}, nil
}

// sliceToMapById converts a slice of user profiles to a map of user profiles by ID
func SliceToMapById(profiles []userProfile) map[string]userProfile {
	profilesById := make(map[string]userProfile, len(profiles))
	for _, profile := range profiles {
		profilesById[profile.ID] = profile
	}
	return profilesById
}
