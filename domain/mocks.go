package domain

var (
	userProfiles = []userProfile{
		{ID: "e3f6fbd6-5bac-11ee-8c99-0242ac120002", Name: "John Doe", Email: "john@example.com"},
		{ID: "102dae5a-5baf-11ee-8c99-0242ac120002", Name: "Jane Smith", Email: "jane@example.com"},
	}

	userRoles = []userRole{
		{UserID: "e3f6fbd6-5bac-11ee-8c99-0242ac120002", Role: "admin"},
		{UserID: "102dae5a-5baf-11ee-8c99-0242ac120002", Role: "user"},
	}
)
