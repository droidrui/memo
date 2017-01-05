package errcode

var Err map[int]string

const (
	Success             = 0
	ServerError         = 1
	UserNotExist        = 2
	PasswordWrong       = 3
	ParamInvalid        = 4
	UserRegistered      = 5
	AccessTokenInvalid  = 6
	AccessTokenTimeout  = 7
	APIInDevelopment    = 8
	RefreshTokenInvalid = 9
	RefreshTokenTimeout = 10
)

func init() {
	Err = make(map[int]string)
	Err[Success] = "success"
	Err[ServerError] = "server error"
	Err[UserNotExist] = "user not exist"
	Err[PasswordWrong] = "password wrong"
	Err[ParamInvalid] = "param invalid"
	Err[UserRegistered] = "user already registered"
	Err[AccessTokenInvalid] = "access token invalid"
	Err[AccessTokenTimeout] = "access token timeout"
	Err[APIInDevelopment] = "api in development"
	Err[RefreshTokenInvalid] = "refresh token invalid"
	Err[RefreshTokenTimeout] = "refresh token timeout"
}
