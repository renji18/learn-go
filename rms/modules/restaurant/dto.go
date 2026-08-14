package restaurant

type AssignOwnerDto struct {
	OwnerId      int `json:"owner_id"`
	RestaurantId int `json:"restaurant_id"`
}

func (l AssignOwnerDto) IsEmpty() bool {
	return l.OwnerId == 0 || l.RestaurantId == 0
}
