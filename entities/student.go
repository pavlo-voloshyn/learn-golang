package entities

type Student struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	CockSize int    `json:"cocksize"`
}
