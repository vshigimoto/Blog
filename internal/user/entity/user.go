package entity

type User struct {
	Id      int    `json:"Id,omitempty"`
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
	About   string `json:"About"`
}
