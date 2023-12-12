package entities

type (
	InsertUsersDto struct {
		Id        string    `json:"id"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		Salt      string    `json:"salt"`
		Email     string    `json:"email"`
		Role      string    `json:"role"`
	}

	Users struct {
		Id        string    `json:"id"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		Salt      string    `json:"salt"`
		Email     string    `json:"email"`
		Role      string    `json:"role"`
		Is_Active bool      `json:"is_active"`
	}

)
