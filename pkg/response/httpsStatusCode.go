package response

const (
	CodeSuccess         = 20001 // Success
	CodeFail            = 40000 // FAIL
	ErrCodeParamInvalid = 20003 // ErrCodeParamInvalid
	ErrInvalidToken     = 30001 // Token is Invalid
	ErrInvalidOtp       = 30002
	ErrSendEmailOtp     = 30003
	// Register code
	ErrCodeUserHassExits = 50001 // User has already been assigned
	// Login code
	ErrCodeOtpNotExists     = 60009 // User does not exist yet
	ErrCodeUserOtpNotExists = 60008

	// User Authentication
	ErrCodeAuthFailed = 40005

	// Two Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed  = 80001
	ErrCodeTwoFactorAuthVerifyFailed = 80002
)

var msg = map[int]string{
	CodeSuccess:          "Success",
	CodeFail:             "Failure",
	ErrCodeParamInvalid:  "ErrCodeParamInvalid",
	ErrInvalidToken:      "Token is invalid",
	ErrCodeUserHassExits: "User has already been assigned",
	ErrInvalidOtp:        "Otp is error",
	ErrSendEmailOtp:      "Send email failed",

	ErrCodeOtpNotExists:     "Otp exists but not registered",
	ErrCodeUserOtpNotExists: "User OTP not exists",

	ErrCodeAuthFailed: "Authentication failed",

	// Two Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed:  "Two Factor Authentication setup failed",
	ErrCodeTwoFactorAuthVerifyFailed: "Two Factor Authentication verify failed",
}
