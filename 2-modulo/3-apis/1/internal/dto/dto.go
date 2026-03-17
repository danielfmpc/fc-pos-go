package dto

type CreateProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutputDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserInputDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutputDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginInputDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutputDto struct {
	AccessToken string `json:"access_token"`
}
