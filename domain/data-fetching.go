package domain

import (
	"context"
	"time"
)

func FetchStaffProfiles(ctx context.Context) map[string]UserProfile {
	// Simulate a 0.5 second wait to fetch data
	time.Sleep(500 * time.Millisecond)

	// Return a mock result
	return userProfiles
}

func FetchUserRolesByID(ctx context.Context, userID string) []string {
	// Simulate a 0.5 second wait to fetch data
	time.Sleep(500 * time.Millisecond)

	// Return a mock result
	return userRoles[userID]
}

func FetchUserPoliciesByID(ctx context.Context, userID string) []string {
	// Simulate a 0.5 second wait to fetch data
	time.Sleep(500 * time.Millisecond)

	// Return a mock result
	return userPolicies[userID]
}

func FetchUserTeamByID(ctx context.Context, userID string) string {
	// Simulate a 0.5 second wait to fetch data
	time.Sleep(500 * time.Millisecond)

	// Return a mock result
	return userTeam[userID]
}
