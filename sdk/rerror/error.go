package rerror

//Error Model
type Error struct {
	Code     int                    `json:"code"`
	Msg      string                 `json:"msg"`
	Trace    string                 `json:"trace,omitempty"`
	DebugMsg string                 `json:"debug_msg,omitempty"`
	Info     map[string]interface{} `json:"info,omitempty"`
}

//ForbiddenErr : forbidden access
func ForbiddenErr(debugMsg ...string) *Error {
	return newError(nil, ForbiddenCode, debugMsg...)
}

//UnauthorizedErr : unauthorized access
func UnauthorizedErr(debugMsg ...string) *Error {
	return newError(nil, UnauthorizedCode, debugMsg...)
}
