package models

type User struct {
	// Id      int    `json:"id" validate:"required" sql:"AUTO_INCREMENT"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type LoginModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Location struct {
	Username  string  `json:"username"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ResponseAll struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}
