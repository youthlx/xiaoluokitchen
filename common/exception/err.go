package exception

type Exception struct {
	ErrCode int
	ErrMsg  string
}

func NewException(code int, msg string) *Exception {
	return &Exception{ErrCode: code, ErrMsg: msg}
}

var (
	EmptyObjException     = NewException(-10001, "object is nil")
	DbConnectException    = NewException(-10002, "database connect fail")
	CacheConnectException = NewException(-10003, "redis connect fail")
	SystemException       = NewException(-10004, "system operate error")
	ParamCheckException   = NewException(-10005, "param check fail")
)
