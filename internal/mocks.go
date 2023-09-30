package internal

var (
	userProfiles = map[string]UserProfile{
		"1": {UserID: "1", Name: "John Doe", JobTitle: "Product Manager"},
		"2": {UserID: "2", Name: "Jane Smith", JobTitle: "Backend Software Enginer"},
		"3": {UserID: "3", Name: "Jack Black", JobTitle: "Frontend Software Enginer"},
	}

	userRoles = map[string][]string{
		"1": {"nonprod_super_admin", "staff_base"},
		"2": {"nonprod_super_admin", "staff_base", "engineer:backend"},
		"3": {"nonprod_super_admin", "staff_base", "engineer:frontend"},
	}
	userPolicies = map[string][]string{
		"1": {"Full Time Employee"},
		"2": {"Backend Engineering", "Full Time Employee"},
		"3": {"Frontend Engineering", "Full Time Employee"},
	}
	userTeam = map[string]string{
		"1": "Finnancial Crime",
		"2": "Core Borrowing",
		"3": "Core Borrowing",
	}
)
