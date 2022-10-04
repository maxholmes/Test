package main

// Member holds of few important details about it.
type Member struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	Last_Name string `json:"last_name"`
	Email     string `json:"email"`
}

// NewMember is Member constructor.
func NewMember(id int, FirstName string, Last_Name string, Email string) Member {
	return Member{id, FirstName, Last_Name, Email}
}
