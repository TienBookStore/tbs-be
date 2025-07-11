package errors

import "errors"

var (
	ErrorEmailExists = errors.New("email đã tồn tại")
	ErrorEmailNotFound = errors.New("email không tồn tại")
	ErrorPasswordIncorrect = errors.New("mật khẩu không đúng")
	ErrorOldPasswordIncorrect = errors.New("mật khẩu cũ không đúng")
	ErrorUserNotFound = errors.New("không tìm thấy người dùng với email này")
	ErrorUserNotActive = errors.New("người dùng chưa được kích hoạt, vui lòng kiểm tra email để kích hoạt tài khoản")
	ErrorInvalidCredentials = errors.New("tên đăng nhập hoặc mật khẩu không đúng")
	ErrorOTPNotFound        = errors.New("mã OTP không tồn tại hoặc đã hết hạn, vui lòng gửi lại OTP")
	ErrorOTPExpired         = errors.New("mã OTP đã hết hạn, vui lòng gửi lại OTP")
	ErrorOTPIncorrect       = errors.New("mã OTP không đúng, vui lòng kiểm tra lại")
	ErrorUserAlreadyActive  = errors.New("tài khoản đã được kích hoạt trước đó, không cần kích hoạt lại")
	ErrorUserUpdateFailed   = errors.New("cập nhật thông tin người dùng không thành công, vui lòng thử lại")
)
