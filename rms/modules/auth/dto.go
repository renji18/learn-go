package auth

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l LoginDto) IsEmpty() bool {
	return l.Email == "" || l.Password == ""
}

type NewUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u NewUserDto) IsEmpty() bool {
	return u.Name == "" || u.Email == "" || u.Password == ""
}
