package req

type ReqSignup struct {
	FullName string `json:"full_name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
