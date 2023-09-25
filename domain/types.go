package domain

type UserRow struct {
	id         string
	name       string
	email      string
	role       string
	policies   []string
	team       string
	collective string
}

type userProfile struct {
	ID    string
	Name  string
	Email string
}
