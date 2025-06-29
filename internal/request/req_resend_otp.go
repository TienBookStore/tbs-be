package request

type ReqResendOTP struct {
	Email string `json:"email" binding:"required,email"`
}
