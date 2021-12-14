package request

type PasswordUpdate struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"min=8"`
}
