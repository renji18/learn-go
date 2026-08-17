package schema

type User struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	IsAdmin    bool       `json:"is_admin"`
	CreatedAt  string     `json:"created_at"`
	Auth       Auth       `json:"auth"`
	Restaurant Restaurant `json:"restaurant"`
}

type Auth struct {
	Id           string `json:"id,omitempty"`
	Password     string `json:"-,omitempty"`
	IsActive     bool   `json:"is_active,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type Restaurant struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	OwnerId   string `json:"owner_id,omitempty"`
	Count     struct {
		RTable   int `json:"tables,omitempty"`
		MenuItem int `json:"menu_item,omitempty"`
	} `json:"_count"`
}

type MenuItem struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Available    bool    `json:"available"`
	CreatedAt    string  `json:"created_at"`
	RestaurantId string  `json:"restaurant_id"`
}

type RTable struct {
	ID           string `json:"id"`
	TableNumber  int    `json:"table_number"`
	ChairCount   int    `json:"chair_count"`
	Occupied     bool   `json:"occupied"`
	Reserved     bool   `json:"reserved"`
	RestaurantId string `json:"restaurant_id,omitempty"`
}

func (r Restaurant) IsEmpty() bool {
	return r.Name == ""
}

func (m MenuItem) IsEmpty() bool {
	return m.Name == "" || m.Category == ""
}

func (rt RTable) IsEmpty() bool {
	return rt.TableNumber == 0 || rt.ChairCount == 0
}
