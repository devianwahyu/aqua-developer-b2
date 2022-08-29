package model

type Customer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Error struct {
	Message string `json:"message"`
}

type Success struct {
	Message string `json:"message"`
}
