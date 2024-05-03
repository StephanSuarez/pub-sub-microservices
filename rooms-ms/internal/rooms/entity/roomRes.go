package entity

type RoomRes struct {
	ID          string
	Name        string
	Description string
	CreatedBy   string
	NumMaxUsers string
	Status      string
	Users       []string
}
