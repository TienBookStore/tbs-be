package request

type ReqSignUp struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqVerifyOTP struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required,len=6"`
}

type ReqLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type ReqResendOTP struct {
	Email string `json:"email" binding:"required,email"`
}

type ReqForgotPassword struct {
	Email string `json:"email" binding:"required,email"`
}

type ReqVerifyForgotPassword struct {
	Email string `json: "email" binding: "required,email"`
	Code  string `json: "otp" binding: "required,len=6"`
}

type ReqResetPassword struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
