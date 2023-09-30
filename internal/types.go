package internal

type UserRow struct {
	ID       string
	name     string
	jobTitle string
	roles    []string
	policies []string
	team     string
}

type UserProfile struct {
	UserID   string
	Name     string
	JobTitle string
}
