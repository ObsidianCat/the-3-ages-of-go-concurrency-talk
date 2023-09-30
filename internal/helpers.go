package internal

import "context"

func NewUserRow(ctx context.Context, profile UserProfile) UserRow {
	// 1 Fetch data
	roles := FetchUserRolesByID(ctx, profile.UserID)
	policies := FetchUserPoliciesByID(ctx, profile.UserID)
	team := FetchUserTeamByID(ctx, profile.UserID)

	// 2 Create user row
	row := UserRow{
		ID:       profile.UserID,
		name:     profile.Name,
		jobTitle: profile.JobTitle,
		roles:    roles,
		policies: policies,
		team:     team,
	}

	return row
}
