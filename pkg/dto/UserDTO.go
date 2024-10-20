package dto

type (
	InputServerUserDTO struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
		Pass  string `json:"pass" validate:"required"`
	}

	OutPutServerUserDTO struct {
		Id    string `json:"_id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
