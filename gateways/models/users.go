package models

type (
	UsersResponse struct {
		Page       int     `json:"page"`
		PerPage    int     `json:"per_page"`
		Total      int     `json:"total"`
		TotalPages int     `json:"total_pages"`
		Data       []Users `json:"data"`
	}

	Users struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	}
)
